package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) RepeatUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "enter":
			switch m.cursor {
			case 0:
				m.state = m.err.ToState
				return m, nil
			case 1:
				switch m.err.ToState {
				case "reg_input_login", "log_input_login":
					return m, tea.Quit
				case "texts":
					m.state = "texts"
					m.TextsList()
				case "files":
					m.state = "files"
					m.FilesList()
				case "cards":
					m.state = "cards"
					m.CardsList()
				default:
					m.state = "menu"
					m.choices = m.currentChoices[m.state]
				}
				m.cursor = 0
				return m, nil
			}
		}
	}

	return m, nil
}

func (m model) RepeatView() string {
	s := "Во время выполнения операций произошла ошибка:\n"
	s += m.err.Err + "\n"
	s += "Повторить операцию?\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	return s
}
