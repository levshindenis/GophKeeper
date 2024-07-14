package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (m *model) AddText() {
	var at []models.Text
	arr := strings.Split(m.helpStr, "///")
	at = append(at, models.Text{Name: arr[0], Description: arr[1], Comment: arr[2], Favourite: false})

	if err := m.db.AddTexts(m.userId, at); err != nil {
		log.Fatalf(err.Error())
	}
	m.state = "texts"
	m.helpStr = ""
	m.TextsList()
}

func (m *model) AddCard() {
	var at []models.Card
	arr := strings.Split(m.helpStr, "///")
	month, err := strconv.Atoi(arr[2])
	if err != nil {
		log.Fatalf(err.Error())
	}
	year, err := strconv.Atoi(arr[3])
	if err != nil {
		log.Fatalf(err.Error())
	}
	cvv, err := strconv.Atoi(arr[4])
	if err != nil {
		log.Fatalf(err.Error())
	}
	at = append(at, models.Card{Bank: arr[0], Number: arr[1], Month: month, Year: year, CVV: cvv,
		Owner: arr[5], Comment: arr[6], Favourite: false})

	if err = m.db.AddCards(m.userId, at); err != nil {
		log.Fatalf(err.Error())
	}
	m.state = "cards"
	m.helpStr = ""
	m.CardsList()
}

func (m *model) AddFile() {
	var at []models.File

	arr := strings.Split(m.helpStr, "///")

	at = append(at, models.File{Name: filepath.Base(arr[0]), Comment: arr[1], Favourite: false})
	if err := m.db.AddFiles(m.userId, at); err != nil {
		log.Fatalf(err.Error())
	}

	login, err := m.db.GetLogin(m.userId)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err = m.cloud.AddFile(login, arr[0]); err != nil {
		log.Fatalf(err.Error())
	}
	source, err := os.Open(arr[0])
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer source.Close()

	destination, err := os.Create("/tmp/keeper/files/" + login + "/" + filepath.Base(arr[0]))
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		log.Fatalf(err.Error())
	}

	m.state = "files"
	m.helpStr = ""
	m.FilesList()
}
