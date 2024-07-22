package main

import (
	"bytes"
	"encoding/json"
	"github.com/levshindenis/GophKeeper/internal/app/models"
	"log"
	"strings"
)

func (m *model) ToServer(localTime string) {
	var (
		allTexts  []models.Text
		allFiles  []models.File
		allCards  []models.Card
		jsonItems []byte
	)

	_, err := m.client.R().Get("http://localhost:8080" + "/user/clear-data")
	if err != nil {
		log.Fatalf(err.Error())
	}

	allTexts, err = m.db.GetUserTexts(m.userId)
	if err != nil {
		log.Fatalf(err.Error())
	}
	allFiles, err = m.db.GetUserFiles(m.userId)
	if err != nil {
		log.Fatalf(err.Error())
	}
	allCards, err = m.db.GetUserCards(m.userId)
	if err != nil {
		log.Fatalf(err.Error())
	}

	jsonItems, err = json.Marshal(allFiles)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if _, err = m.client.R().SetBody(bytes.NewBuffer(jsonItems)).
		Post("http://localhost:8080" + "/user/add-files"); err != nil {
		log.Fatalf(err.Error())
	}

	jsonItems, err = json.Marshal(allTexts)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if _, err = m.client.R().SetBody(bytes.NewBuffer(jsonItems)).
		Post("http://localhost:8080" + "/user/add-texts"); err != nil {
		log.Fatalf(err.Error())
	}

	jsonItems, err = json.Marshal(allCards)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if _, err = m.client.R().SetBody(bytes.NewBuffer(jsonItems)).
		Post("http://localhost:8080" + "/user/add-cards"); err != nil {
		log.Fatalf(err.Error())
	}

	if _, err = m.client.R().SetBody(strings.NewReader(localTime)).
		Post("http://localhost:8080" + "/user/set-time"); err != nil {
		log.Fatalf(err.Error())
	}
}
