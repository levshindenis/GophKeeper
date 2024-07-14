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

func (m *model) ChangeText() {
	var (
		at []models.ChText
	)
	newName := m.textItem.Name
	newDescription := m.textItem.Description
	newComment := m.textItem.Comment
	arr := strings.Split(m.helpStr, "///")
	if arr[1] != "" {
		newName = arr[1]
	}
	if arr[2] != "" {
		newDescription = arr[2]
	}
	if arr[3] != "" {
		newComment = arr[3]
	}

	at = append(at, models.ChText{OldName: m.textItem.Name, NewName: newName,
		NewDescription: newDescription, NewComment: newComment, NewFavourite: false})

	if err := m.db.ChangeTexts(m.userId, at); err != nil {
		log.Fatalf(err.Error())
	}

	m.state = "text_view"
	m.helpStr = newName
	m.cursor = 5
	m.TextInfo()
}

func (m *model) ChangeCard() {
	var (
		at []models.ChCard
	)
	newBank := m.cardItem.Bank
	newNumber := m.cardItem.Number
	newMonth := m.cardItem.Month
	newYear := m.cardItem.Year
	newCVV := m.cardItem.CVV
	newOwner := m.cardItem.Owner
	newComment := m.cardItem.Comment
	arr := strings.Split(m.helpStr, "///")
	if arr[1] != "" {
		newBank = arr[1]
	}
	if arr[2] != "" {
		newNumber = arr[2]
	}
	if arr[3] != "" {
		helpPer, _ := strconv.Atoi(arr[3])
		newMonth = helpPer
	}
	if arr[4] != "" {
		helpPer, _ := strconv.Atoi(arr[4])
		newYear = helpPer
	}
	if arr[5] != "" {
		helpPer, _ := strconv.Atoi(arr[5])
		newCVV = helpPer
	}
	if arr[6] != "" {
		newOwner = arr[6]
	}
	if arr[7] != "" {
		newComment = arr[7]
	}

	at = append(at, models.ChCard{OldBank: m.cardItem.Bank, OldNumber: m.cardItem.Number, NewBank: newBank,
		NewNumber: newNumber, NewMonth: newMonth, NewYear: newYear, NewCVV: newCVV, NewOwner: newOwner,
		NewComment: newComment, NewFavourite: false})

	if err := m.db.ChangeCards(m.userId, at); err != nil {
		log.Fatalf(err.Error())
	}

	m.state = "card_view"
	m.helpStr = newBank + " " + newNumber
	m.cursor = 8
	m.CardInfo()
}

func (m *model) ChangeFile() {
	var (
		at []models.ChFile
	)
	newName := m.fileItem.Name
	newComment := m.fileItem.Comment
	arr := strings.Split(m.helpStr, "///")
	if arr[1] != "" {
		newName = arr[1]
	}
	if arr[2] != "" {
		newComment = arr[2]
	}

	at = append(at, models.ChFile{OldName: m.fileItem.Name, NewName: filepath.Base(newName),
		NewComment: newComment, NewFavourite: false})

	if err := m.db.ChangeFiles(m.userId, at); err != nil {
		log.Fatalf(err.Error())
	}

	if newName != m.fileItem.Name {
		login, err := m.db.GetLogin(m.userId)
		if err != nil {
			log.Fatalf(err.Error())
		}
		if err = m.cloud.DeleteFile(login, arr[0]); err != nil {
			log.Fatalf(err.Error())
		}
		if err = m.cloud.AddFile(login, newName); err != nil {
			log.Fatalf(err.Error())
		}

		if err = os.Remove("/tmp/keeper/files/" + login + "/" + arr[0]); err != nil {
			log.Fatalf(err.Error())
		}

		source, err := os.Open(newName)
		if err != nil {
			log.Fatalf(err.Error())
		}
		defer source.Close()

		destination, err := os.Create("/tmp/keeper/files/" + login + "/" + filepath.Base(newName))
		if err != nil {
			log.Fatalf(err.Error())
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}

	m.state = "file_view"
	m.helpStr = filepath.Base(newName)
	m.cursor = 5
	m.FileInfo()
}
