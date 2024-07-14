package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) FavouritesUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			m.state = "menu"
			m.choices = m.currentChoices[m.state]
			m.cursor = 0
			return m, nil
		}
	}

	return m, nil
}
