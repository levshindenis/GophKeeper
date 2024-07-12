package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (s *Server) ChangeTextsH() {
	var (
		str     string
		items   []models.ChText
		oldItem models.Text
		newItem models.Text
	)

	fmt.Println("Введите название записи:   ")
	fmt.Scanf("%s\n", &str)

	resp, err := s.client.R().Get(s.address + "/user/texts/" + str)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(resp.Body(), &oldItem); err != nil {
		panic(err)
	}

	fmt.Println("\nName: ", oldItem.Name)
	fmt.Println("Description: ", oldItem.Description)
	fmt.Println("Comment: ", oldItem.Comment)
	fmt.Println("Favourite: ", oldItem.Favourite)

	fmt.Println("\nВведите новое название записи:   ")
	fmt.Scanf("%s\n", &str)
	if str != "" {
		newItem.Name = str
	}

	fmt.Println("Введите новый текст записи:   ")
	fmt.Scanf("%s\n", &str)
	if str != "" {
		newItem.Description = str
	}

	fmt.Println("Введите новый комментарий записи:  ")
	fmt.Scanf("%s\n", &str)
	if str != "" {
		newItem.Comment = str
	}

	items = append(items, models.ChText{OldName: oldItem.Name, OldDescription: oldItem.Description,
		OldComment: oldItem.Comment, OldFavourite: oldItem.Favourite, NewName: newItem.Name,
		NewDescription: newItem.Description, NewComment: newItem.Comment, NewFavourite: false})

	marsh, err := json.Marshal(items)
	if err != nil {
		panic(err)
	}

	resp, err = s.client.R().SetBody(bytes.NewReader(marsh)).Post(s.address + "/user/change-texts")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Status())

	if err = s.f.Event(context.Background(), "mainpage"); err != nil {
		panic(err)
	}
}
