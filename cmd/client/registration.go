package main

import (
	"log"
	"strings"

	"github.com/levshindenis/GophKeeper/internal/app/models"
	"github.com/levshindenis/GophKeeper/internal/app/tools"
)

func (m *model) Registration() {
	arr := strings.Split(m.helpStr, "///")
	reg := models.Register{Login: arr[0], Password: arr[1], Word: arr[2]}
	flag, userId, err := m.db.AddUser(reg)

	if flag || err != nil {
		if flag {
			m.err.Err = "Логин уже занят. Выберите другой!"
		} else {
			m.err.Err = err.Error()
		}
		m.state = "repeat"
		m.err.ToState = "reg_input_login"
		m.userId = ""
	} else {
		m.state = "menu"
		m.userId = userId
	}
	m.choices = m.currentChoices[m.state]
	m.helpStr = ""
	if m.state == "menu" {
		if err = tools.MakeFilesDirectory(m.userId); err != nil {
			log.Fatalf(err.Error())
		}
	}
}
