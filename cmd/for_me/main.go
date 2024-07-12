// Package client используется для отправки запросов со стороны клиента. Испульзуется finite state machine.
package main

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/looplab/fsm"
)

// Server - основная структура для хранения значений сервера и клиента
type Server struct {
	client  *resty.Client
	cookie  string
	address string
	choice  string
	m       map[string]string
	f       *fsm.FSM
}

// NewServer - функция для создания нового Server
func NewServer() *Server {
	client := resty.New()
	m := map[string]string{
		"1":  "reg",
		"2":  "log",
		"3":  "logout",
		"4":  "deleteAccount",
		"5":  "addTexts",
		"7":  "addCards",
		"8":  "changeTexts",
		"10": "changeCards",
		"11": "deleteTexts",
		"13": "deleteCards",
		"14": "listTexts",
		"16": "listCards",
		"20": "exit",
	}
	return &Server{
		client:  client,
		cookie:  "",
		address: "http://localhost:8080",
		choice:  "",
		m:       m,
	}
}

func main() {
	server := NewServer()
	server.f = fsm.NewFSM(
		"zero",
		fsm.Events{
			{Name: "go", Src: []string{"zero"}, Dst: "main"},
			{Name: "reg", Src: []string{"main"}, Dst: "regH"},
			{Name: "log", Src: []string{"main"}, Dst: "logH"},
			{Name: "logout", Src: []string{"main"}, Dst: "logoutH"},
			{Name: "exit", Src: []string{"main"}, Dst: "ExitH"},
			{Name: "deleteAccount", Src: []string{"main"}, Dst: "deleteAccountH"},
			{Name: "addTexts", Src: []string{"main"}, Dst: "addTextsH"},
			{Name: "changeTexts", Src: []string{"main"}, Dst: "changeTextsH"},
			{Name: "deleteTexts", Src: []string{"main"}, Dst: "deleteTextsH"},
			{Name: "listTexts", Src: []string{"main"}, Dst: "listTextsH"},
			{Name: "addCards", Src: []string{"main"}, Dst: "addCardsH"},
			{Name: "changeCards", Src: []string{"main"}, Dst: "changeCardsH"},
			{Name: "deleteCards", Src: []string{"main"}, Dst: "deleteCardsH"},
			{Name: "listCards", Src: []string{"main"}, Dst: "listCardsH"},
			{Name: "mainpage",
				Src: []string{"regH", "logH", "logoutH", "deleteAccountH", "addTextsH", "deleteTextsH", "changeTextsH",
					"listTextsH", "addCardsH", "deleteCardsH", "changeCardsH", "listCardsH"},
				Dst: "main"},
		},
		fsm.Callbacks{
			"main":          func(_ context.Context, _ *fsm.Event) { server.SelectAction() },
			"reg":           func(_ context.Context, _ *fsm.Event) { server.RegH() },
			"log":           func(_ context.Context, _ *fsm.Event) { server.LogH() },
			"logout":        func(_ context.Context, _ *fsm.Event) { server.LogoutH() },
			"exit":          func(_ context.Context, _ *fsm.Event) { server.ExitH() },
			"deleteAccount": func(_ context.Context, _ *fsm.Event) { server.DeleteAccountH() },
			"addTexts":      func(_ context.Context, _ *fsm.Event) { server.AddTextsH() },
			"changeTexts":   func(_ context.Context, _ *fsm.Event) { server.ChangeTextsH() },
			"deleteTexts":   func(_ context.Context, _ *fsm.Event) { server.DeleteTextsH() },
			"listTexts":     func(_ context.Context, _ *fsm.Event) { server.ListTextsH() },
			"addCards":      func(_ context.Context, _ *fsm.Event) { server.AddCardsH() },
			"changeCards":   func(_ context.Context, _ *fsm.Event) { server.ChangeCardsH() },
			"deleteCards":   func(_ context.Context, _ *fsm.Event) { server.DeleteCardsH() },
			"listCards":     func(_ context.Context, _ *fsm.Event) { server.ListCardsH() },
		},
	)

	if err := server.f.Event(context.Background(), "go"); err != nil {
		panic(err)
	}
}
