package main

import (
	"encoding/json"
	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (m *model) ToLocal() {
	var (
		allTexts []models.Text
		allFiles []models.File
		allCards []models.Card
	)

	resp, err := m.client.R().Get("http://localhost:8080" + "/user/all-texts")
	if resp.StatusCode() != 200 || err != nil {
		m.ErrorState(string(resp.Body()), "menu")
		if err != nil {
			m.err.Err = err.Error()
		}
		return
	}

	if err = json.Unmarshal(resp.Body(), &allTexts); err != nil {
		m.ErrorState(err.Error(), "menu")
		return
	}

	resp, err = m.client.R().Get("http://localhost:8080" + "/user/all-files")
	if err != nil {
		m.ErrorState(err.Error(), "menu")
		return
	}
	if err = json.Unmarshal(resp.Body(), &allFiles); err != nil {
		m.ErrorState(err.Error(), "menu")
		return
	}

	resp, err = m.client.R().Get("http://localhost:8080" + "/user/all-cards")
	if err != nil {
		m.ErrorState(err.Error(), "menu")
		return
	}
	if err = json.Unmarshal(resp.Body(), &allCards); err != nil {
		m.ErrorState(err.Error(), "menu")
		return
	}

	if err = m.db.AddTexts(m.userId, allTexts); err != nil {
		m.ErrorState(err.Error(), "menu")
		return
	}

	if err = m.db.AddFiles(m.userId, allFiles); err != nil {
		m.ErrorState(err.Error(), "menu")
		return
	}

	if err = m.db.AddCards(m.userId, allCards); err != nil {
		m.ErrorState(err.Error(), "menu")
		return
	}
}
