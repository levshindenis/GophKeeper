package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/levshindenis/GophKeeper/internal/app/models"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/levshindenis/GophKeeper/internal/app/tools"
)

func (m *model) Registration() {
	arr := strings.Split(m.helpStr, "///")

	user := models.Register{Login: arr[0], Password: arr[1], Word: arr[2]}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		log.Fatalf(err.Error())
	}
	resp, err := m.client.R().SetBody(bytes.NewBuffer(jsonUser)).Post("http://localhost:8080" + "/register")
	if err != nil {
		log.Fatalf(err.Error())
	}

	if resp.StatusCode() != 201 {
		m.err.Err = string(resp.Body())
		m.state = "repeat"
		m.err.ToState = "reg_input_login"
		m.userId = ""
	} else {
		m.state = "menu"
		m.userId = arr[0]
	}

	m.choices = m.currentChoices[m.state]
	m.helpStr = ""

	if m.state == "menu" {
		if err = m.cloud.CreateBucket(m.userId); err != nil {
			log.Fatalf(err.Error())
		}
		if err = tools.MakeFilesDirectory(m.userId); err != nil {
			log.Fatalf(err.Error())
		}

		cryptoBytes, err := tools.GenerateCrypto(12)
		if err != nil {
			log.Fatalf(err.Error())
		}
		m.secretKey = fmt.Sprintf("%x", cryptoBytes)

		f, err := os.OpenFile("../../.env", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			log.Fatalf(err.Error())
		}
		if _, err = f.WriteString(strings.ToUpper(m.userId) + "_LOCAL=" + m.secretKey + "\n"); err != nil {
			log.Fatalf(err.Error())
		}
		f.Close()

		resp, err = m.client.R().Get("http://localhost:8080" + "/user/update-time")
		if err != nil {
			log.Fatalf(err.Error())
		}
		if err = m.db.AddUpdateTime(m.userId, string(resp.Body())); err != nil {
			log.Fatalf(err.Error())
		}

		if err = godotenv.Load("../../.env"); err != nil {
			log.Fatal("Error loading .env file")
		}

		m.secretKey = os.Getenv(strings.ToUpper(m.userId) + "_LOCAL")
	}
}
