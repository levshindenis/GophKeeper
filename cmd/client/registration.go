package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/levshindenis/GophKeeper/internal/app/models"
	"github.com/levshindenis/GophKeeper/internal/app/tools"
)

func (m *model) Registration() {
	arr := strings.Split(m.helpStr, "///")

	user := models.Register{Login: arr[0], Password: arr[1], Word: arr[2]}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		m.ErrorState(err.Error(), "reg_input_login")
		return
	}

	m.state = "menu"
	m.userId = arr[0]

	resp, err := m.client.R().SetBody(bytes.NewBuffer(jsonUser)).Post("http://localhost:8080" + "/register")
	if resp.StatusCode() != 201 || err != nil {
		m.ErrorState(string(resp.Body()), "reg_input_login")
		if err != nil {
			m.err.Err = err.Error()
		}
		return
	}

	m.choices = m.currentChoices[m.state]
	m.helpStr = ""

	if m.state == "menu" {
		if err = m.cloud.CreateBucket(m.userId); err != nil {
			m.ErrorState(err.Error(), "reg_input_login")
			return
		}
		if err = tools.MakeFilesDirectory(m.userId); err != nil {
			m.ErrorState(err.Error(), "reg_input_login")
			return
		}

		cryptoBytes, err := tools.GenerateCrypto(12)
		if err != nil {
			m.ErrorState(err.Error(), "reg_input_login")
			return
		}
		m.secretKey = fmt.Sprintf("%x", cryptoBytes)

		f, err := os.OpenFile("../../.env", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			log.Fatalf(err.Error())
		}
		if _, err = f.WriteString(strings.ToUpper(m.userId) + "_LOCAL=" + m.secretKey + "\n"); err != nil {
			m.ErrorState(err.Error(), "reg_input_login")
			return
		}
		f.Close()

		resp, err = m.client.R().Get("http://localhost:8080" + "/user/update-time")
		if resp.StatusCode() != 200 || err != nil {
			m.ErrorState(string(resp.Body()), "reg_input_login")
			if err != nil {
				m.err.Err = err.Error()
			}
			return
		}

		if err = m.db.AddUpdateTime(m.userId, string(resp.Body())); err != nil {
			m.ErrorState(err.Error(), "reg_input_login")
			return
		}

		if err = godotenv.Load("../../.env"); err != nil {
			m.ErrorState(err.Error(), "reg_input_login")
			return
		}

		m.secretKey = os.Getenv(strings.ToUpper(m.userId) + "_LOCAL")
	}
}
