package main

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (m *model) ToServer(localTime string) {
	var (
		allTexts  []models.Text
		allFiles  []models.File
		allCards  []models.Card
		jsonItems []byte
	)

	resp, err := m.client.R().Get("http://localhost:8080" + "/user/clear-data")
	if resp.StatusCode() != 200 || err != nil {
		m.ErrorState(string(resp.Body()), "menu")
		if err != nil {
			m.err.Err = err.Error()
		}
		return
	}

	allTexts, err = m.db.GetUserTexts(m.userId)
	if err != nil {
		m.ErrorState(err.Error(), "menu")
		return
	}
	allFiles, err = m.db.GetUserFiles(m.userId)
	if err != nil {
		m.ErrorState(err.Error(), "menu")
		return
	}
	allCards, err = m.db.GetUserCards(m.userId)
	if err != nil {
		m.ErrorState(err.Error(), "menu")
		return
	}

	jsonItems, err = json.Marshal(allFiles)
	if err != nil {
		m.ErrorState(err.Error(), "menu")
		return
	}
	if _, err = m.client.R().SetBody(bytes.NewBuffer(jsonItems)).
		Post("http://localhost:8080" + "/user/add-files"); err != nil {
		m.ErrorState(err.Error(), "menu")
		return
	}

	jsonItems, err = json.Marshal(allTexts)
	if err != nil {
		m.ErrorState(err.Error(), "menu")
		return
	}
	if _, err = m.client.R().SetBody(bytes.NewBuffer(jsonItems)).
		Post("http://localhost:8080" + "/user/add-texts"); err != nil {
		m.ErrorState(err.Error(), "menu")
		return
	}

	jsonItems, err = json.Marshal(allCards)
	if err != nil {
		m.ErrorState(err.Error(), "menu")
		return
	}
	if _, err = m.client.R().SetBody(bytes.NewBuffer(jsonItems)).
		Post("http://localhost:8080" + "/user/add-cards"); err != nil {
		m.ErrorState(err.Error(), "menu")
		return
	}

	if _, err = m.client.R().SetBody(strings.NewReader(localTime)).
		Post("http://localhost:8080" + "/user/set-time"); err != nil {
		m.ErrorState(err.Error(), "menu")
		return
	}
}
