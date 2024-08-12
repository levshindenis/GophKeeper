package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/levshindenis/GophKeeper/internal/app/models"
	"github.com/levshindenis/GophKeeper/internal/app/tools"
)

func (m *model) AddText() {
	var at []models.Text
	arr := strings.Split(m.helpStr, "///")
	at = append(at, models.Text{Name: tools.Encrypt(arr[0], m.secretKey),
		Description: tools.Encrypt(arr[1], m.secretKey), Comment: tools.Encrypt(arr[2], m.secretKey),
		Favourite: "Нет"})

	if err := m.db.AddTexts(m.userId, at); err != nil {
		m.ErrorState(err.Error(), "add_text_name")
		return
	}
	if err := m.db.SetUpdateTime(m.userId, time.Now().Format(time.RFC3339)); err != nil {
		m.ErrorState(err.Error(), "add_text_name")
		return
	}
	m.state = "texts"
	m.helpStr = ""
	m.TextsList()
}

func (m *model) AddCard() {
	var at []models.Card
	arr := strings.Split(m.helpStr, "///")

	at = append(at, models.Card{Bank: tools.Encrypt(arr[0], m.secretKey), Number: tools.Encrypt(arr[1], m.secretKey),
		Month: tools.Encrypt(arr[2], m.secretKey), Year: tools.Encrypt(arr[3], m.secretKey),
		CVV: tools.Encrypt(arr[4], m.secretKey), Owner: tools.Encrypt(arr[5], m.secretKey),
		Comment: tools.Encrypt(arr[6], m.secretKey), Favourite: "Нет"})

	if err := m.db.AddCards(m.userId, at); err != nil {
		m.ErrorState(err.Error(), "add_card_name")
		return
	}
	if err := m.db.SetUpdateTime(m.userId, time.Now().Format(time.RFC3339)); err != nil {
		m.ErrorState(err.Error(), "add_card_name")
		return
	}
	m.state = "cards"
	m.helpStr = ""
	m.CardsList()
}

func (m *model) AddFile() {
	var at []models.File

	arr := strings.Split(m.helpStr, "///")

	at = append(at, models.File{Name: tools.Encrypt(filepath.Base(arr[0]), m.secretKey),
		Comment: tools.Encrypt(arr[1], m.secretKey), Favourite: "Нет"})
	if err := m.db.AddFiles(m.userId, at); err != nil {
		m.ErrorState(err.Error(), "add_file_name")
		return
	}

	if err := m.cloud.AddFile(m.userId, arr[0]); err != nil {
		m.ErrorState(err.Error(), "add_file_name")
		return
	}
	if err := m.db.SetUpdateTime(m.userId, time.Now().Format(time.RFC3339)); err != nil {
		m.ErrorState(err.Error(), "add_file_name")
		return
	}
	source, err := os.Open(arr[0])
	if err != nil {
		m.ErrorState(err.Error(), "add_file_name")
		return
	}
	defer source.Close()

	destination, err := os.Create("/tmp/keeper/files/" + m.userId + "/" + filepath.Base(arr[0]))
	if err != nil {
		m.ErrorState(err.Error(), "add_file_name")
		return
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		m.ErrorState(err.Error(), "add_file_name")
		return
	}

	m.state = "files"
	m.helpStr = ""
	m.FilesList()
}
