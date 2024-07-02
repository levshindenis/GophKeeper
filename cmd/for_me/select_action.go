package main

import (
	"context"
	"fmt"
	"strconv"
)

func (s *Server) SelectAction() {
	fmt.Println("Cookie: ", s.cookie)
	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1) Регистрация")
		fmt.Println("2) Вход")
		fmt.Println("3) Выход")
		fmt.Println("4) Удалить аккаунт")
		fmt.Println("5) Добавить текст")
		fmt.Println("6) Добавить файл")
		fmt.Println("7) Добавить карту")
		fmt.Println("8) Поменять текст")
		fmt.Println("9) Поменять файл")
		fmt.Println("10) Поменять карту")
		fmt.Println("11) Удалить текст")
		fmt.Println("12) Удалить файл")
		fmt.Println("13) Удалить карту")
		fmt.Println("14) Посмотреть тексты")
		fmt.Println("15) Посмотреть файлы")
		fmt.Println("16) Посмотреть карты")
		fmt.Println("17) Посмотреть избранное")
		fmt.Println("18) Exit")
		fmt.Println("===========================")
		fmt.Print("Ввод:    ")
		fmt.Scanf("%s", &s.choice)
		fmt.Println(s.choice)
		if number, _ := strconv.Atoi(s.choice); 0 < number && number < 19 {
			break
		}
		fmt.Println("Bad answer. Please repeat!")
	}

	if err := s.f.Event(context.Background(), s.m[s.choice]); err != nil {
		panic(err)
	}
}
