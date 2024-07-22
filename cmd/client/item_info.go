package main

import (
	"github.com/levshindenis/GophKeeper/internal/app/models"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gen2brain/dlgs"

	"github.com/levshindenis/GophKeeper/internal/app/tools"
)

func (m *model) TextInfo() {
	item, err := m.db.GetText(m.userId, m.helpStr)
	if err != nil {
		log.Printf(err.Error())
	}
	m.textItem = item
	arr := []string{"Название записи: " + tools.Decrypt(item.Name, m.secretKey),
		"Содержание записи: " + tools.Decrypt(item.Description, m.secretKey),
		"Комментарий к записи: " + tools.Decrypt(item.Comment, m.secretKey),
		"В избранном: " + item.Favourite}
	m.choices = append(arr, "", "Изменить запись", "Добавить в избранное", "Удалить запись", "Назад")
	if item.Favourite == "Да" {
		m.choices = append(arr, "", "Изменить запись", "Удалить из избранного", "Удалить запись", "Назад")
	}
}

func (m *model) CardInfo() {
	item, err := m.db.GetCard(m.userId, m.helpStr)
	if err != nil {
		log.Printf(err.Error())
	}
	m.cardItem = item
	arr := []string{
		"Банк: " + tools.Decrypt(item.Bank, m.secretKey),
		"Номер карты: " + tools.Decrypt(item.Number, m.secretKey),
		"Дата окончания действия карты: " + tools.Decrypt(item.Month, m.secretKey) + "/" + tools.Decrypt(item.Year, m.secretKey),
		"СVV: " + tools.Decrypt(item.CVV, m.secretKey),
		"Владелец карты: " + tools.Decrypt(item.Owner, m.secretKey),
		"Комментарий к карте: " + tools.Decrypt(item.Comment, m.secretKey),
		"В избранном: " + item.Favourite}
	m.choices = append(arr, "", "Изменить запись", "Добавить в избранное", "Удалить запись", "Назад")
	if item.Favourite == "Да" {
		m.choices = append(arr, "", "Изменить запись", "Удалить из избранного", "Удалить запись", "Назад")
	}
}

func (m *model) FileInfo() {
	item, err := m.db.GetFile(m.userId, m.helpStr)
	if err != nil {
		log.Printf(err.Error())
	}
	m.fileItem = item
	arr := []string{
		"Название файла: " + tools.Decrypt(item.Name, m.secretKey),
		"Комментарий к файлу: " + tools.Decrypt(item.Comment, m.secretKey),
		"В избранном: " + item.Favourite}
	m.choices = append(arr, "", "Изменить запись", "Добавить в избранное", "Удалить запись", "Назад")
	if item.Favourite == "Да" {
		m.choices = append(arr, "", "Изменить запись", "Удалить из избранного", "Удалить запись", "Назад")
	}
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
			switch m.cursor {
			case len(m.choices) - 4:
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
			case len(m.choices) - 3:
				switch {
				case strings.Contains(m.state, "text"):
					arr := models.ChText{OldName: m.textItem.Name, NewName: m.textItem.Name,
						NewDescription: m.textItem.Description, NewComment: m.textItem.Comment,
						NewFavourite: "Да"}
					if m.textItem.Favourite == "Да" {
						arr = models.ChText{OldName: m.textItem.Name, NewName: m.textItem.Name,
							NewDescription: m.textItem.Description, NewComment: m.textItem.Comment,
							NewFavourite: "Нет"}
					}
					if err := m.db.ChangeTexts(m.userId, arr); err != nil {
						log.Fatalf(err.Error())
					}
					m.TextInfo()
					m.cursor = 5
				case strings.Contains(m.state, "file"):
					arr := models.ChFile{OldName: m.fileItem.Name, NewName: m.fileItem.Name,
						NewComment: m.fileItem.Comment, NewFavourite: "Да"}
					if m.fileItem.Favourite == "Да" {
						arr = models.ChFile{OldName: m.fileItem.Name, NewName: m.fileItem.Name,
							NewComment: m.fileItem.Comment, NewFavourite: "Нет"}
					}
					if err := m.db.ChangeFiles(m.userId, arr); err != nil {
						log.Fatalf(err.Error())
					}
					m.FileInfo()
					m.cursor = 4
				case strings.Contains(m.state, "card"):
					arr := models.ChCard{OldBank: m.cardItem.Bank, OldNumber: m.cardItem.Number,
						NewBank: m.cardItem.Bank, NewNumber: m.cardItem.Number, NewMonth: m.cardItem.Month,
						NewYear: m.cardItem.Year, NewCVV: m.cardItem.CVV, NewOwner: m.cardItem.Owner,
						NewComment: m.cardItem.Comment, NewFavourite: "Да"}
					if m.cardItem.Favourite == "Да" {
						arr = models.ChCard{OldBank: m.cardItem.Bank, OldNumber: m.cardItem.Number,
							NewBank: m.cardItem.Bank, NewNumber: m.cardItem.Number, NewMonth: m.cardItem.Month,
							NewYear: m.cardItem.Year, NewCVV: m.cardItem.CVV, NewOwner: m.cardItem.Owner,
							NewComment: m.cardItem.Comment, NewFavourite: "Нет"}
					}
					if err := m.db.ChangeCard(m.userId, arr); err != nil {
						log.Fatalf(err.Error())
					}
					m.CardInfo()
					m.cursor = 8
				}
				return m, nil
			case len(m.choices) - 2:
				switch {
				case strings.Contains(m.state, "text"):
					if err := m.db.DeleteTexts(m.userId, []string{m.helpStr}, ""); err != nil {
						log.Fatalf(err.Error())
					}
					m.TextsList()
				case strings.Contains(m.state, "file"):
					if err := m.db.DeleteFiles(m.userId, []string{m.helpStr}, ""); err != nil {
						log.Fatalf(err.Error())
					}
					if err := m.cloud.DeleteFile(m.userId, "/tmp/keeper/files/"+m.userId+"/"+m.helpStr); err != nil {
						log.Fatalf(err.Error())
					}
					if err := os.Remove("/tmp/keeper/files/" + m.userId + "/" + m.helpStr); err != nil {
						log.Fatalf(err.Error())
					}
					m.FilesList()
				case strings.Contains(m.state, "card"):
					if err := m.db.DeleteCards(m.userId, []string{m.helpStr}, ""); err != nil {
						log.Fatalf(err.Error())
					}
					m.CardsList()
				}
			case len(m.choices) - 1:
				switch {
				case strings.Contains(m.state, "text"):
					m.TextsList()
				case strings.Contains(m.state, "file"):
					m.FilesList()
				case strings.Contains(m.state, "card"):
					m.CardsList()
				}
			}

			m.state = strings.Split(m.state, "_")[0] + "s"
			m.cursor = 0
			m.helpStr = ""
			return m, nil
		}
	}
	return m, nil
}
