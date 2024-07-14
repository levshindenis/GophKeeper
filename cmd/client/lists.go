package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gen2brain/dlgs"
	"github.com/skratchdot/open-golang/open"
	"log"
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

func (m *model) FavouritesList() {
	arr, err := m.db.ListFavourites(m.userId)
	if err != nil {
		log.Printf(err.Error())
	}
	m.choices = append(arr, "", "Назад")
	m.cursor = len(m.choices) - 1
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
				if m.state == "add_file_name" {
					filePath, flag, err := dlgs.File("Выберите файл для загрузки:", "", false)
					if err != nil {
						log.Fatalf(err.Error())
					}
					if !flag {
						m.state = "files"
						return m, nil
					}
					m.helpStr = filePath + "///"
					m.state = "add_file_comment"
				}
				return m, nil
			}
			if m.cursor == len(m.choices)-1 {
				m.cursor = 0
				m.state = "menu"
				m.choices = m.currentChoices[m.state]
				return m, nil
			}
			infoType := m.state[:len(m.state)-1]
			m.state = infoType + "_view"
			m.helpStr = m.choices[m.cursor]
			if infoType == "text" {
				m.cursor = 5
				m.TextInfo()
			}
			if infoType == "file" {
				m.cursor = 4
				login, err := m.db.GetLogin(m.userId)
				if err != nil {
					log.Fatalf(err.Error())
				}

				filePath := "/tmp/keeper/files/" + login + "/" + m.helpStr
				if err = open.Run(filePath); err != nil {
					log.Printf(err.Error())
				}

				if err = m.cloud.DeleteFile(login, m.helpStr); err != nil {
					log.Printf(err.Error())
				}
				if err = m.cloud.AddFile(login, filePath); err != nil {
					log.Printf(err.Error())
				}
				m.FileInfo()
			}
			if infoType == "card" {
				m.cursor = 8
				m.CardInfo()
			}
			return m, nil
		}
	}
	return m, nil
}
