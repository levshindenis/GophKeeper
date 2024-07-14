package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func (m model) MenuUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.cursor == 0 {
				m.state = "texts"
				m.TextsList()
			}
			if m.cursor == 1 {
				m.state = "files"
				m.FilesList()
			}
			if m.cursor == 2 {
				m.state = "cards"
				m.CardsList()
			}
			if m.cursor == 3 {
				m.state = "favourites"
				m.FavouritesList()
				return m, nil
			}
			if m.cursor == 4 {
				m.state = "start"
				m.choices = m.currentChoices[m.state]
			}
			if m.cursor == 5 {
				if err := m.db.Close(); err != nil {
					log.Fatalf(err.Error())
				}
				return m, tea.Quit
			}

			m.helpStr = ""
			m.cursor = 0
			return m, nil
		}
	}
	return m, nil
}
