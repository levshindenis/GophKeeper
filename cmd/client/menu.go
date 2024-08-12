package main

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
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
			switch m.cursor {
			case 0:
				m.state = "texts"
				m.TextsList()
			case 1:
				m.state = "files"
				m.FilesList()
			case 2:
				m.state = "cards"
				m.CardsList()
			case 3:
				m.state = "favourites"
				m.FavouritesList()
				return m, nil
			case 4:
				var (
					localTime  string
					serverTime string
				)

				m.state = "start"
				m.choices = m.currentChoices[m.state]

				resp, err := m.client.R().Get("http://localhost:8080" + "/user/update-time")
				if resp.StatusCode() != 200 || err != nil {
					m.ErrorState(string(resp.Body()), "menu")
					if err != nil {
						m.err.Err = err.Error()
					}
					return m, nil
				}
				serverTime = string(resp.Body())

				localTime, err = m.db.GetUpdateTime(m.userId)
				if err != nil {
					m.ErrorState(err.Error(), "menu")
					return m, nil
				}

				if localTime > serverTime {
					m.ToServer(localTime)
				}

				if _, err = m.client.R().Get("http://localhost:8080" + "/user/logout"); err != nil {
					m.ErrorState(err.Error(), "menu")
					return m, nil
				}
			case 5:
				m.state = "start"
				m.choices = m.currentChoices[m.state]

				if err := m.db.DeleteAccount(m.userId, "local"); err != nil {
					m.ErrorState(err.Error(), "menu")
					return m, nil
				}
				if _, err := m.client.R().Get("http://localhost:8080" + "/user/delete-account"); err != nil {
					m.ErrorState(err.Error(), "menu")
					return m, nil
				}
				if err := os.RemoveAll("/tmp/keeper/files/" + m.userId); err != nil {
					m.ErrorState(err.Error(), "menu")
					return m, nil
				}
				if err := m.cloud.DeleteAccount(m.userId); err != nil {
					m.ErrorState(err.Error(), "menu")
					return m, nil
				}
			case 6:
				var (
					localTime  string
					serverTime string
				)

				resp, err := m.client.R().Get("http://localhost:8080" + "/user/update-time")
				if resp.StatusCode() != 200 || err != nil {
					m.ErrorState(string(resp.Body()), "menu")
					if err != nil {
						m.err.Err = err.Error()
					}
					return m, nil
				}

				serverTime = string(resp.Body())

				localTime, err = m.db.GetUpdateTime(m.userId)
				if err != nil {
					m.ErrorState(err.Error(), "menu")
					return m, nil
				}

				if localTime > serverTime {
					m.ToServer(localTime)
				}

				if err = m.db.Close(); err != nil {
					m.ErrorState(err.Error(), "menu")
					return m, nil
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
