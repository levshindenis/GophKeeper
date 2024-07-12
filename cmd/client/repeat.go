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
			if m.cursor == 0 {
				m.state = m.err.ToState
				return m, nil
			}
			m.cursor = 0
			return m, tea.Quit
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
