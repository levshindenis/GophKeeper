package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) InputUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.helpStr += m.textInput.Value() + "///"
			m.textInput.Reset()
			m.state = m.nextState[m.state]
			if m.state == "registration" {
				m.Registration()
			}
			if m.state == "login" {
				m.Login()
			}
			if m.state == "add_text" {
				m.AddText()
			}
			if m.state == "add_card" {
				m.AddCard()
			}
			return m, nil
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) InputView() string {
	if m.state == "reg_input_login" || m.state == "log_input_login" {
		return fmt.Sprintf("Введите логин:\n%s\n", m.textInput.View())
	}
	if m.state == "reg_input_password" || m.state == "log_input_password" {
		return fmt.Sprintf("Введите пароль:\n%s\n", m.textInput.View())
	}
	if m.state == "reg_input_word" {
		return fmt.Sprintf("Введите ключевое слово:\n%s\n", m.textInput.View())
	}
	if m.state == "add_text_name" {
		return fmt.Sprintf("Введите название записи:\n%s\n", m.textInput.View())
	}
	if m.state == "add_text_description" {
		return fmt.Sprintf("Введите текст записи:\n%s\n", m.textInput.View())
	}
	if m.state == "add_text_comment" {
		return fmt.Sprintf("Введите комментарий к записи:\n%s\n", m.textInput.View())
	}
	if m.state == "add_card_name" {
		return fmt.Sprintf("Введите название банка:\n%s\n", m.textInput.View())
	}
	if m.state == "add_card_number" {
		return fmt.Sprintf("Введите номер карты:\n%s\n", m.textInput.View())
	}
	if m.state == "add_card_month" {
		return fmt.Sprintf("Введите месяц окончания дествия карты:\n%s\n", m.textInput.View())
	}
	if m.state == "add_card_year" {
		return fmt.Sprintf("Введите код окончания действия карты:\n%s\n", m.textInput.View())
	}
	if m.state == "add_card_cvv" {
		return fmt.Sprintf("Введите cvv код:\n%s\n", m.textInput.View())
	}
	if m.state == "add_card_owner" {
		return fmt.Sprintf("Введите данные владельца карты:\n%s\n", m.textInput.View())
	}
	if m.state == "add_card_comment" {
		return fmt.Sprintf("Введите комментарий к карте:\n%s\n", m.textInput.View())
	}
	return ""
}
