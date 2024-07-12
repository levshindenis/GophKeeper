package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (s *Server) AddCardsH() {
	var (
		bank, number, owner, comment string
		year, month, cvv             int
		items                        []models.Card
	)

	fmt.Println("Введите название банка:   ")
	fmt.Scanf("%s\n", &bank)
	fmt.Println("Введите номер карты:  ")
	fmt.Scanf("%s\n", &number)
	fmt.Println("Введите месяц окончания действия карты:  ")
	fmt.Scanf("%d\n", &month)
	fmt.Println("Введите год окончания действия карты:   ")
	fmt.Scanf("%d\n", &year)
	fmt.Println("Введите CVV код:   ")
	fmt.Scanf("%d\n", &cvv)
	fmt.Println("Введите имя и фамилию владельца карты:   ")
	fmt.Scanf("%s\n", &owner)
	fmt.Println("Введите комментарий:   ")
	fmt.Scanf("%s\n", &comment)

	items = append(items, models.Card{Bank: bank, Number: number, Month: month, Year: year, CVV: cvv, Owner: owner,
		Comment: comment, Favourite: false})

	jsonItems, err := json.Marshal(items)
	if err != nil {
		panic(err)
	}

	resp, err := s.client.R().SetBody(bytes.NewBuffer(jsonItems)).Post(s.address + "/user/add-cards")
	if err != nil {
		panic(err)
	}

	fmt.Println("Ответ:")

	fmt.Println(resp.Status())
	fmt.Println(resp.String())

	if err = s.f.Event(context.Background(), "mainpage"); err != nil {
		panic(err)
	}
}
