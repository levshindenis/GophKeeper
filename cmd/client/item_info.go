package main

import (
	"log"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
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
	m.choices = append(arr, "", "Изменить запись", "Удалить запись", "Назад")
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
	m.choices = append(arr, "", "Изменить запись", "Удалить запись", "Назад")
}

func (m model) ItemUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.cursor > len(m.choices)-3 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			if m.cursor == len(m.choices)-3 {
				m.helpStr += "///"
				m.state = "change_" + strings.Split(m.state, "_")[0] + "_name"
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
