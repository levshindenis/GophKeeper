package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gen2brain/dlgs"
	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (m *model) TextInfo() {
	item, err := m.db.GetText(m.userId, m.helpStr)
	if err != nil {
		log.Printf(err.Error())
	}
	m.textItem = item
	arr := []string{"Название записи: " + item.Name,
		"Содержание записи: " + item.Description,
		"Комментарий к записи: " + item.Comment,
		"В избранном: " + strconv.FormatBool(item.Favourite)}
	m.choices = append(arr, "", "Изменить запись", "Добавить в избранное", "Удалить запись", "Назад")
}

func (m *model) CardInfo() {
	item, err := m.db.GetCard(m.userId, m.helpStr)
	if err != nil {
		log.Printf(err.Error())
	}
	m.cardItem = item
	arr := []string{
		"Банк: " + item.Bank,
		"Номер карты: " + item.Number,
		"Дата окончания действия карты: " + strconv.Itoa(item.Month) + "/" + strconv.Itoa(item.Year),
		"СVV: " + strconv.Itoa(item.CVV),
		"Владелец карты: " + item.Owner,
		"Комментарий к карте: " + item.Comment,
		"В избранном: " + strconv.FormatBool(item.Favourite)}
	m.choices = append(arr, "", "Изменить данные карты", "Добавить в избранное", "Удалить карту", "Назад")
}

func (m *model) FileInfo() {
	item, err := m.db.GetFile(m.userId, m.helpStr)
	if err != nil {
		log.Printf(err.Error())
	}
	m.fileItem = item
	arr := []string{
		"Название файла: " + item.Name,
		"Комментарий к файлу: " + item.Comment,
		"В избранном: " + strconv.FormatBool(item.Favourite)}
	m.choices = append(arr, "", "Заменить файл", "Добавить в избранное", "Удалить файл", "Назад")
}

func (m model) ItemUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.cursor > len(m.choices)-4 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			if m.cursor == len(m.choices)-4 {
				m.helpStr += "///"
				m.state = "change_" + strings.Split(m.state, "_")[0] + "_name"
				if m.state == "change_file_name" {
					filePath, _, err := dlgs.File("Выберите файл для замены:", "", false)
					if err != nil {
						log.Fatalf(err.Error())
					}
					m.helpStr += filePath + "///"
				}
				return m, nil
			}

			if m.cursor == len(m.choices)-3 {
				if strings.Contains(m.state, "text") {
					var arr []models.ChText
					arr = append(arr, models.ChText{OldName: m.textItem.Name, NewName: m.textItem.Name,
						NewDescription: m.textItem.Description, NewComment: m.textItem.Comment, NewFavourite: true})
					if err := m.db.ChangeTexts(m.userId, arr); err != nil {
						log.Fatalf(err.Error())
					}
					m.TextInfo()
					m.cursor = 5
				}
				if strings.Contains(m.state, "file") {
					var arr []models.ChFile
					arr = append(arr, models.ChFile{OldName: m.fileItem.Name, NewName: m.fileItem.Name,
						NewComment: m.fileItem.Comment, NewFavourite: true})
					if err := m.db.ChangeFiles(m.userId, arr); err != nil {
						log.Fatalf(err.Error())
					}
					m.FileInfo()
					m.cursor = 4
				}
				if strings.Contains(m.state, "card") {
					var arr []models.ChCard
					arr = append(arr, models.ChCard{OldBank: m.cardItem.Bank, OldNumber: m.cardItem.Number,
						NewBank: m.cardItem.Bank, NewNumber: m.cardItem.Number, NewMonth: m.cardItem.Month,
						NewYear: m.cardItem.Year, NewCVV: m.cardItem.CVV, NewOwner: m.cardItem.Owner,
						NewComment: m.cardItem.Comment, NewFavourite: true})
					if err := m.db.ChangeCards(m.userId, arr); err != nil {
						log.Fatalf(err.Error())
					}
					m.CardInfo()
					m.cursor = 8
				}
				return m, nil
			}

			m.state = strings.Split(m.state, "_")[0] + "s"

			if m.cursor == len(m.choices)-2 { // удалить
				if m.state == "texts" {
					if err := m.db.DeleteTexts(m.userId, []string{m.helpStr}); err != nil {
						log.Fatalf(err.Error())
					}
					m.TextsList()
				} else if m.state == "files" {
					if err := m.db.DeleteFiles(m.userId, []string{m.helpStr}); err != nil {
						log.Fatalf(err.Error())
					}
					login, err := m.db.GetLogin(m.userId)
					if err != nil {
						log.Fatalf(err.Error())
					}
					if err = m.cloud.DeleteFile(login, "/tmp/keeper/files/"+login+"/"+m.helpStr); err != nil {
						log.Fatalf(err.Error())
					}
					if err = os.Remove("/tmp/keeper/files/" + login + "/" + m.helpStr); err != nil {
						log.Fatalf(err.Error())
					}
					m.FilesList()
				} else {
					if err := m.db.DeleteCards(m.userId, []string{m.helpStr}); err != nil {
						log.Fatalf(err.Error())
					}
					m.CardsList()
				}
			}
			if m.cursor == len(m.choices)-1 {
				if m.state == "texts" {
					m.TextsList()
				} else if m.state == "files" {
					m.FilesList()
				} else {
					m.CardsList()
				}
			}
			m.cursor = 0
			m.helpStr = ""
			return m, nil
		}
	}

	return m, nil
}
