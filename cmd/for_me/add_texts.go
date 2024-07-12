package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (s *Server) AddTextsH() {
	var (
		name, description, comment string
		items                      []models.Text
	)
	fmt.Println("Введите название записи:   ")
	fmt.Scanf("%s\n", &name)
	fmt.Println("Введите текст записи:  ")
	fmt.Scanf("%s\n", &description)
	fmt.Println("Введите комментарий к записи:  ")
	fmt.Scanf("%s\n", &comment)

	items = append(items, models.Text{Name: name, Description: description, Comment: comment, Favourite: false})

	jsonItems, err := json.Marshal(items)
	if err != nil {
		panic(err)
	}

	resp, err := s.client.R().SetBody(bytes.NewBuffer(jsonItems)).Post(s.address + "/user/add-texts")
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
