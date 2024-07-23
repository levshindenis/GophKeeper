package main

import (
	"path"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gen2brain/dlgs"
	"github.com/skratchdot/open-golang/open"

	"github.com/levshindenis/GophKeeper/internal/app/tools"
)

func (m *model) TextsList() {
	m.encChoices = []string{}
	arr, err := m.db.ListTexts(m.userId)
	if err != nil {
		m.ErrorState(err.Error(), "main")
		return
	}
	for i := range arr {
		m.encChoices = append(m.encChoices, arr[i])
		arr[i] = tools.Decrypt(arr[i], m.secretKey)
	}
	m.choices = append(arr, "Добавить запись", "Назад")
}

func (m *model) FilesList() {
	m.encChoices = []string{}
	arr, err := m.db.ListFiles(m.userId)
	if err != nil {
		m.ErrorState(err.Error(), "main")
		return
	}
	for i := range arr {
		m.encChoices = append(m.encChoices, arr[i])
		arr[i] = tools.Decrypt(arr[i], m.secretKey)
	}
	m.choices = append(arr, "Добавить файл", "Назад")
}

func (m *model) CardsList() {
	m.encChoices = []string{}
	arr, err := m.db.ListCards(m.userId)
	if err != nil {
		m.ErrorState(err.Error(), "main")
		return
	}
	for i := range arr {
		m.encChoices = append(m.encChoices, arr[i])
		arr[i] = tools.Decrypt(strings.Split(arr[i], "///")[0], m.secretKey) + " " +
			tools.Decrypt(strings.Split(arr[i], "///")[1], m.secretKey)
	}
	m.choices = append(arr, "Добавить карту", "Назад")
}

func (m *model) FavouritesList() {
	arr, err := m.db.ListFavourites(m.userId)
	if err != nil {
		m.ErrorState(err.Error(), "main")
		return
	}
	for i := range arr {
		switch {
		case arr[i] == "Cards:" || arr[i] == "Files:" || arr[i] == "Texts:":
			continue
		case len(strings.Split(arr[i], "///")) > 1:
			arr[i] = "    " + tools.Decrypt(strings.Split(arr[i], "///")[0], m.secretKey) + " " +
				tools.Decrypt(strings.Split(arr[i], "///")[1], m.secretKey)
		default:
			arr[i] = "    " + tools.Decrypt(arr[i], m.secretKey)
		}
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
		case "enter":
			switch m.cursor {
			case len(m.choices) - 2:
				m.cursor = 0
				m.state = "add_" + m.state[:len(m.state)-1] + "_name"
				if m.state == "add_file_name" {
					filePath, flag, err := dlgs.File("Выберите файл для загрузки:", "", false)
					if err != nil {
						m.ErrorState(err.Error(), "files")
					}
					if !flag {
						m.state = "files"
						return m, nil
					}
					m.helpStr = filePath + "///"
					m.state = "add_file_comment"
				}
				return m, nil
			case len(m.choices) - 1:
				m.cursor = 0
				m.state = "menu"
				m.choices = m.currentChoices[m.state]
				return m, nil
			}

			infoType := m.state[:len(m.state)-1]
			m.state = infoType + "_view"
			m.helpStr = m.encChoices[m.cursor]

			switch infoType {
			case "text":
				m.cursor = 5
				m.TextInfo()
			case "file":
				m.cursor = 4

				filePath := path.Join("/tmp/keeper/files", m.userId, tools.Decrypt(m.helpStr, m.secretKey))
				if err := open.Run(filePath); err != nil {

					if err1 := m.cloud.GetFile(m.userId, tools.Decrypt(m.helpStr, m.secretKey), "/tmp/keeper/files"); err1 != nil {
						m.ErrorState(err1.Error(), "files")
						return m, nil
					}
					if err1 := open.Run(filePath); err1 != nil {
						m.ErrorState(err1.Error(), "files")
						return m, nil
					}
				}

				if err := m.cloud.DeleteFile(m.userId, tools.Decrypt(m.helpStr, m.secretKey)); err != nil {
					m.ErrorState(err.Error(), "files")
					return m, nil
				}
				if err := m.cloud.AddFile(m.userId, filePath); err != nil {
					m.ErrorState(err.Error(), "files")
					return m, nil
				}
				m.FileInfo()
			case "card":
				m.cursor = 8
				m.CardInfo()
			}
			return m, nil
		}
	}
	return m, nil
}
