package main

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/levshindenis/GophKeeper/internal/app/models"
	"github.com/levshindenis/GophKeeper/internal/app/tools"
)

func (m *model) Login() {
	arr := strings.Split(m.helpStr, "///")
	user := models.Login{Login: arr[0], Password: arr[1]}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		m.ErrorState(err.Error(), "log_input_login")
		return
	}

	m.state = "menu"
	m.userId = user.Login

	resp, err := m.client.R().SetBody(bytes.NewBuffer(jsonUser)).Post("http://localhost:8080" + "/login")
	if resp.StatusCode() != 200 || err != nil {
		m.ErrorState(string(resp.Body()), "log_input_login")
		if err != nil {
			m.err.Err = err.Error()
		}
		return
	}

	m.helpStr = ""
	m.choices = m.currentChoices[m.state]
	if m.state == "menu" {
		var (
			serverTime string
		)

		if err = tools.MakeFilesDirectory(m.userId); err != nil {
			m.ErrorState(err.Error(), "log_input_login")
			return
		}

		resp, err = m.client.R().Get("http://localhost:8080" + "/user/update-time")
		if resp.StatusCode() != 200 || err != nil {
			m.ErrorState(string(resp.Body()), "log_input_login")
			if err != nil {
				m.err.Err = err.Error()
			}
			return
		}

		serverTime = string(resp.Body())

		localTime, err1 := m.db.GetUpdateTime(m.userId)
		if err1 != nil {
			m.ToLocal()
			if err2 := m.db.AddUpdateTime(m.userId, serverTime); err2 != nil {
				m.ErrorState(err2.Error(), "log_input_login")
				return
			}
		} else {
			if localTime > serverTime {
				m.ToServer(localTime)
			}
			if localTime < serverTime {
				if err2 := m.db.DeleteTexts(m.userId, nil, "all"); err2 != nil {
					m.ErrorState(err2.Error(), "log_input_login")
					return
				}
				if err2 := m.db.DeleteFiles(m.userId, nil, "all"); err2 != nil {
					m.ErrorState(err2.Error(), "log_input_login")
					return
				}
				if err2 := m.db.DeleteCards(m.userId, nil, "all"); err2 != nil {
					m.ErrorState(err2.Error(), "log_input_login")
					return
				}
				m.ToLocal()
				if err2 := m.db.SetUpdateTime(m.userId, serverTime); err2 != nil {
					m.ErrorState(err2.Error(), "log_input_login")
					return
				}
			}
		}

		if err = godotenv.Load("../../.env"); err != nil {
			m.ErrorState(err.Error(), "log_input_login")
			return
		}

		m.secretKey = os.Getenv(strings.ToUpper(m.userId) + "_LOCAL")
	}
}
