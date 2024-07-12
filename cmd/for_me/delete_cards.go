package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

func (s *Server) DeleteCardsH() {
	var (
		name  string
		items []string
	)
	fmt.Println("Введите название записи:   ")
	fmt.Scanf("%s\n", &name)

	items = append(items, name)

	jsonItems, err := json.Marshal(items)
	if err != nil {
		panic(err)
	}

	resp, err := s.client.R().SetBody(bytes.NewBuffer(jsonItems)).Post(s.address + "/user/delete-cards")
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
