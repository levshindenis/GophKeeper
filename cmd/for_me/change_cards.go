package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (s *Server) ChangeCardsH() {
	var (
		str     string
		value   int
		items   []models.ChCard
		oldItem models.Card
		newItem models.Card
	)

	fmt.Println("Введите название записи:   ")
	fmt.Scanf("%s\n", &str)

	resp, err := s.client.R().Get(s.address + "/user/cards/" + str)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(resp.Body(), &oldItem); err != nil {
		panic(err)
	}

	fmt.Println("\nName: ", oldItem.Bank)
	fmt.Println("Description: ", oldItem.Number)
	fmt.Println("Comment: ", oldItem.Month)
	fmt.Println("Favourite: ", oldItem.Year)
	fmt.Println("\nName: ", oldItem.CVV)
	fmt.Println("Description: ", oldItem.Owner)
	fmt.Println("Comment: ", oldItem.Comment)

	fmt.Println("\nВведите новое название банка:   ")
	fmt.Scanf("%s\n", &str)
	if str != "" {
		newItem.Bank = str
	}

	fmt.Println("Введите новый номер карты:   ")
	fmt.Scanf("%s\n", &str)
	if str != "" {
		newItem.Number = str
	}

	fmt.Println("Введите новый месяц:  ")
	fmt.Scanf("%d\n", &value)
	if str != "" {
		newItem.Month = value
	}

	fmt.Println("\nВведите новый год:   ")
	fmt.Scanf("%d\n", &value)
	if str != "" {
		newItem.Year = value
	}
	fmt.Println("\nВведите новый код:   ")
	fmt.Scanf("%d\n", &value)
	if str != "" {
		newItem.CVV = value
	}
	fmt.Println("\nВведите нового владельца:   ")
	fmt.Scanf("%s\n", &str)
	if str != "" {
		newItem.Owner = str
	}
	fmt.Println("\nВведите новый комментарий:   ")
	fmt.Scanf("%s\n", &str)
	if str != "" {
		newItem.Comment = str
	}

	items = append(items, models.ChCard{OldBank: oldItem.Bank, OldNumber: oldItem.Number, OldMonth: oldItem.Month,
		OldYear: oldItem.Year, OldCVV: oldItem.CVV, OldOwner: oldItem.Owner, OldComment: oldItem.Comment,
		OldFavourite: oldItem.Favourite, NewBank: newItem.Bank, NewNumber: newItem.Number, NewMonth: newItem.Month,
		NewYear: newItem.Year, NewCVV: newItem.CVV, NewOwner: newItem.Owner, NewComment: newItem.Comment,
		NewFavourite: false})

	marsh, err := json.Marshal(items)
	if err != nil {
		panic(err)
	}

	resp, err = s.client.R().SetBody(bytes.NewReader(marsh)).Post(s.address + "/user/change-cards")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Status())

	if err = s.f.Event(context.Background(), "mainpage"); err != nil {
		panic(err)
	}
}
