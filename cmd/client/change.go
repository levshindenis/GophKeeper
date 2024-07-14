package main

import (
	"log"
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
	m.cursor = 9
	m.CardInfo()
}
