package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *model) TextsList() {
	arr, err := m.db.ListTexts(m.userId)
	if err != nil {
		log.Printf(err.Error())
	}
	m.choices = append(arr, "Добавить запись", "Назад")
}

func (m *model) FilesList() {
	arr, err := m.db.ListFiles(m.userId)
	if err != nil {
		log.Printf(err.Error())
	}
	m.choices = append(arr, "Добавить файл", "Назад")
}

func (m *model) CardsList() {
	arr, err := m.db.ListCards(m.userId)
	if err != nil {
		log.Printf(err.Error())
	}
	m.choices = append(arr, "Добавить карту", "Назад")
}

func (m model) ListsUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			if m.cursor == len(m.choices)-2 {
				m.cursor = 0
				m.state = "add_" + m.state[:len(m.state)-1] + "_name"
				return m, nil
			}
			if m.cursor == len(m.choices)-1 {
				m.cursor = 0
				m.state = "menu"
				m.choices = m.currentChoices[m.state]
				return m, nil
			}
		}
	}
	return m, nil
}
