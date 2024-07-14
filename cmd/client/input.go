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
			if m.state == "add_file" {
				m.AddFile()
			}
			if m.state == "change_text" {
				m.ChangeText()
			}
			if m.state == "change_card" {
				m.ChangeCard()
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
	if m.state == "add_file_comment" {
		return fmt.Sprintf("Введите комментарий к файлу:\n%s\n", m.textInput.View())
	}
	if m.state == "change_text_name" {
		return fmt.Sprintf("Введите новое название записи: (если хотите оставить текущее, нажмите enter)\n%s%s\n%s\n",
			"Текущее название записи: ", m.textItem.Name, m.textInput.View())
	}
	if m.state == "change_text_description" {
		return fmt.Sprintf("Введите новое содержание записи: (если хотите оставить текущее, нажмите enter)\n%s%s\n%s\n",
			"Текущее содержание записи: ", m.textItem.Description, m.textInput.View())
	}
	if m.state == "change_text_comment" {
		return fmt.Sprintf("Введите новый комментарий к записи: (если хотите оставить текущий, нажмите enter)\n%s%s\n%s\n",
			"Текущий комментарий: ", m.textItem.Comment, m.textInput.View())
	}
	if m.state == "change_card_name" {
		return fmt.Sprintf("Введите новое название банка: (если хотите оставить текущее, нажмите enter)\n%s%s\n%s\n",
			"Текущее название банка: ", m.cardItem.Bank, m.textInput.View())
	}
	if m.state == "change_card_number" {
		return fmt.Sprintf("Введите новый номер карты: (если хотите оставить текущий, нажмите enter)\n%s%s\n%s\n",
			"Текущий номер: ", m.cardItem.Number, m.textInput.View())
	}
	if m.state == "change_card_month" {
		return fmt.Sprintf("Введите новый месяц окончания действия карты: (если хотите оставить текущий, нажмите enter)\n%s%d\n%s\n",
			"Текущий месяц: ", m.cardItem.Month, m.textInput.View())
	}
	if m.state == "change_card_year" {
		return fmt.Sprintf("Введите новый год окончания действия карты: (если хотите оставить текущий, нажмите enter)\n%s%d\n%s\n",
			"Текущий год: ", m.cardItem.Year, m.textInput.View())
	}
	if m.state == "change_card_cvv" {
		return fmt.Sprintf("Введите новый cvv код: (если хотите оставить текущий, нажмите enter)\n%s%d\n%s\n",
			"Текущий cvv код: ", m.cardItem.CVV, m.textInput.View())
	}
	if m.state == "change_card_owner" {
		return fmt.Sprintf("Введите нового владельца карты: (если хотите оставить текущего, нажмите enter)\n%s%s\n%s\n",
			"Текущий владелец: ", m.cardItem.Owner, m.textInput.View())
	}
	if m.state == "change_card_comment" {
		return fmt.Sprintf("Введите новый комментарий: (если хотите оставить текущий, нажмите enter)\n%s%s\n%s\n",
			"Текущий комментарий: ", m.cardItem.Comment, m.textInput.View())
	}
	return ""
}
