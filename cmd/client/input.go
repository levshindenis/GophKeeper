package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/levshindenis/GophKeeper/internal/app/tools"
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
			switch m.state {
			case "registration":
				m.Registration()
			case "login":
				m.Login()
			case "add_text":
				m.AddText()
			case "add_card":
				m.AddCard()
			case "add_file":
				m.AddFile()
			case "change_text":
				m.ChangeText()
			case "change_card":
				m.ChangeCard()
			case "change_file":
				m.ChangeFile()
			}
			return m, nil
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) InputView() string {
	switch m.state {
	case "reg_input_login", "log_input_login":
		return fmt.Sprintf("Введите логин:\n%s\n", m.textInput.View())
	case "reg_input_password", "log_input_password":
		return fmt.Sprintf("Введите пароль:\n%s\n", m.textInput.View())
	case "reg_input_word":
		return fmt.Sprintf("Введите ключевое слово:\n%s\n", m.textInput.View())
	case "add_text_name":
		return fmt.Sprintf("Введите название записи:\n%s\n", m.textInput.View())
	case "add_text_description":
		return fmt.Sprintf("Введите текст записи:\n%s\n", m.textInput.View())
	case "add_text_comment":
		return fmt.Sprintf("Введите комментарий к записи:\n%s\n", m.textInput.View())
	case "add_card_name":
		return fmt.Sprintf("Введите название банка:\n%s\n", m.textInput.View())
	case "add_card_number":
		return fmt.Sprintf("Введите номер карты:\n%s\n", m.textInput.View())
	case "add_card_month":
		return fmt.Sprintf("Введите месяц окончания дествия карты:\n%s\n", m.textInput.View())
	case "add_card_year":
		return fmt.Sprintf("Введите код окончания действия карты:\n%s\n", m.textInput.View())
	case "add_card_cvv":
		return fmt.Sprintf("Введите cvv код:\n%s\n", m.textInput.View())
	case "add_card_owner":
		return fmt.Sprintf("Введите данные владельца карты:\n%s\n", m.textInput.View())
	case "add_card_comment":
		return fmt.Sprintf("Введите комментарий к карте:\n%s\n", m.textInput.View())
	case "add_file_comment":
		return fmt.Sprintf("Введите комментарий к файлу:\n%s\n", m.textInput.View())
	case "change_text_name":
		return fmt.Sprintf("Введите новое название записи: (если хотите оставить текущее, нажмите enter)\n%s%s\n%s\n",
			"Текущее название записи: ", tools.Decrypt(m.textItem.Name, m.secretKey), m.textInput.View())
	case "change_text_description":
		return fmt.Sprintf("Введите новое содержание записи: (если хотите оставить текущее, нажмите enter)\n%s%s\n%s\n",
			"Текущее содержание записи: ", tools.Decrypt(m.textItem.Description, m.secretKey), m.textInput.View())
	case "change_text_comment":
		return fmt.Sprintf("Введите новый комментарий к записи: (если хотите оставить текущий, нажмите enter)\n%s%s\n%s\n",
			"Текущий комментарий: ", tools.Decrypt(m.textItem.Comment, m.secretKey), m.textInput.View())
	case "change_card_name":
		return fmt.Sprintf("Введите новое название банка: (если хотите оставить текущее, нажмите enter)\n%s%s\n%s\n",
			"Текущее название банка: ", tools.Decrypt(m.cardItem.Bank, m.secretKey), m.textInput.View())
	case "change_card_number":
		return fmt.Sprintf("Введите новый номер карты: (если хотите оставить текущий, нажмите enter)\n%s%s\n%s\n",
			"Текущий номер: ", tools.Decrypt(m.cardItem.Number, m.secretKey), m.textInput.View())
	case "change_card_month":
		return fmt.Sprintf("Введите новый месяц окончания действия карты: (если хотите оставить текущий, нажмите enter)\n%s%d\n%s\n",
			"Текущий месяц: ", tools.Decrypt(m.cardItem.Month, m.secretKey), m.textInput.View())
	case "change_card_year":
		return fmt.Sprintf("Введите новый год окончания действия карты: (если хотите оставить текущий, нажмите enter)\n%s%d\n%s\n",
			"Текущий год: ", tools.Decrypt(m.cardItem.Year, m.secretKey), m.textInput.View())
	case "change_card_cvv":
		return fmt.Sprintf("Введите новый cvv код: (если хотите оставить текущий, нажмите enter)\n%s%d\n%s\n",
			"Текущий cvv код: ", tools.Decrypt(m.cardItem.CVV, m.secretKey), m.textInput.View())
	case "change_card_owner":
		return fmt.Sprintf("Введите нового владельца карты: (если хотите оставить текущего, нажмите enter)\n%s%s\n%s\n",
			"Текущий владелец: ", tools.Decrypt(m.cardItem.Owner, m.secretKey), m.textInput.View())
	case "change_card_comment":
		return fmt.Sprintf("Введите новый комментарий: (если хотите оставить текущий, нажмите enter)\n%s%s\n%s\n",
			"Текущий комментарий: ", tools.Decrypt(m.cardItem.Comment, m.secretKey), m.textInput.View())
	case "change_file_name":
		return fmt.Sprintf("Введите новый комментарий: (если хотите оставить текущий, нажмите enter)\n%s%s\n%s\n",
			"Текущий комментарий: ", tools.Decrypt(m.fileItem.Comment, m.secretKey), m.textInput.View())
	}
	return ""
}
