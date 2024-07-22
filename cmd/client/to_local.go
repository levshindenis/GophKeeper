package main

import (
	"encoding/json"
	"github.com/levshindenis/GophKeeper/internal/app/models"
	"log"
)

func (m *model) ToLocal() {
	var (
		allTexts []models.Text
		allFiles []models.File
		allCards []models.Card
	)

	resp, err := m.client.R().Get("http://localhost:8080" + "/user/all-texts")
	if err != nil {
		log.Fatalf(err.Error())
	}
	if err = json.Unmarshal(resp.Body(), &allTexts); err != nil {
		log.Fatalf(err.Error())
	}

	resp, err = m.client.R().Get("http://localhost:8080" + "/user/all-files")
	if err != nil {
		log.Fatalf(err.Error())
	}
	if err = json.Unmarshal(resp.Body(), &allFiles); err != nil {
		log.Fatalf(err.Error())
	}

	resp, err = m.client.R().Get("http://localhost:8080" + "/user/all-cards")
	if err != nil {
		log.Fatalf(err.Error())
	}
	if err = json.Unmarshal(resp.Body(), &allCards); err != nil {
		log.Fatalf(err.Error())
	}

	if err = m.db.AddTexts(m.userId, allTexts); err != nil {
		log.Fatalf(err.Error())
	}

	if err = m.db.AddFiles(m.userId, allFiles); err != nil {
		log.Fatalf(err.Error())
	}

	if err = m.db.AddCards(m.userId, allCards); err != nil {
		log.Fatalf(err.Error())
	}
}
