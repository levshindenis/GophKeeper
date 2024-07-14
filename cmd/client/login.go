package main

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	"github.com/levshindenis/GophKeeper/internal/app/models"
	"github.com/levshindenis/GophKeeper/internal/app/tools"
)

func (m *model) Login() {
	arr := strings.Split(m.helpStr, "///")
	mlog := models.Login{Login: arr[0], Password: arr[1]}
	userId, err := m.db.CheckUser(mlog)

	if err != nil {
		m.err.Err = err.Error()
		if errors.Is(err, sql.ErrNoRows) {
			m.err.Err = "Неверный логин или пароль"
		}
		m.state = "repeat"
		m.err.ToState = "log_input_login"
		m.userId = ""
	} else {
		m.state = "menu"
		m.userId = userId
	}
	m.helpStr = ""
	m.choices = m.currentChoices[m.state]
	if m.state == "menu" {

		if err = tools.MakeFilesDirectory(mlog.Login); err != nil {
			log.Fatalf(err.Error())
		}
	}
}
