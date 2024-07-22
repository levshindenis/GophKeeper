package main

import (
	"bytes"
	"encoding/json"
	"github.com/levshindenis/GophKeeper/internal/app/models"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/levshindenis/GophKeeper/internal/app/tools"
)

func (m *model) Login() {
	arr := strings.Split(m.helpStr, "///")
	user := models.Login{Login: arr[0], Password: arr[1]}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	resp, err := m.client.R().SetBody(bytes.NewBuffer(jsonUser)).Post("http://localhost:8080" + "/login")
	if err != nil {
		panic(err)
	}

	m.state = "menu"
	m.userId = user.Login
	if resp.StatusCode() != 200 {
		m.err.Err = string(resp.Body())
		m.state = "repeat"
		m.err.ToState = "log_input_login"
		m.userId = ""
	}

	m.helpStr = ""
	m.choices = m.currentChoices[m.state]
	if m.state == "menu" {
		var (
			serverTime string
		)

		if err = tools.MakeFilesDirectory(m.userId); err != nil {
			log.Fatalf(err.Error())
		}

		resp, err = m.client.R().Get("http://localhost:8080" + "/user/update-time")
		if err == nil {
			serverTime = string(resp.Body())

			localTime, err1 := m.db.GetUpdateTime(m.userId)
			if err1 != nil {
				m.ToLocal()
				if err2 := m.db.AddUpdateTime(m.userId, serverTime); err2 != nil {
					log.Fatalf(err2.Error())
				}
			} else {
				if localTime > serverTime {
					m.ToServer(localTime)
				}
				if localTime < serverTime {
					if err2 := m.db.DeleteTexts(m.userId, nil, "all"); err2 != nil {
						log.Fatalf(err2.Error())
					}
					if err2 := m.db.DeleteFiles(m.userId, nil, "all"); err2 != nil {
						log.Fatalf(err2.Error())
					}
					if err2 := m.db.DeleteCards(m.userId, nil, "all"); err2 != nil {
						log.Fatalf(err2.Error())
					}
					m.ToLocal()
					if err2 := m.db.SetUpdateTime(m.userId, serverTime); err2 != nil {
						log.Fatalf(err2.Error())
					}
				}
			}
		}
		if err = godotenv.Load("../../.env"); err != nil {
			log.Fatal("Error loading .env file")
		}

		m.secretKey = os.Getenv(strings.ToUpper(m.userId) + "_LOCAL")
	}
}
