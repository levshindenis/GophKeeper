package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (s *Server) RegH() {
	var login, password, word string
	fmt.Println("Введите логин:   ")
	fmt.Scanf("%s\n", &login)
	fmt.Println("Введите пароль:  ")
	fmt.Scanf("%s\n", &password)
	fmt.Println("Введите кодовое слово:  ")
	fmt.Scanf("%s\n", &word)

	user := models.Register{Login: login, Password: password, Word: word}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	resp, err := s.client.R().SetBody(bytes.NewBuffer(jsonUser)).Post(s.address + "/register")
	if err != nil {
		panic(err)
	}

	if resp.StatusCode() == 201 {
		s.cookie = resp.Cookies()[0].Value
	}

	fmt.Println("Ответ:")

	fmt.Println(resp.Status())
	fmt.Println(resp.String())

	if err = s.f.Event(context.Background(), "mainpage"); err != nil {
		panic(err)
	}
}
