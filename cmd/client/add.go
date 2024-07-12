package main

import (
	"log"
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
	m.choices = m.currentChoices[m.state]
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
	m.choices = m.currentChoices[m.state]
	m.helpStr = ""
	m.CardsList()
}
