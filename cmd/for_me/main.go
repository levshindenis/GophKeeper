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
		"18": "exit",
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
			{Name: "mainpage",
				Src: []string{"regH", "logH", "logoutH", "deleteAccountH"},
				Dst: "main"},
		},
		fsm.Callbacks{
			"main":          func(_ context.Context, _ *fsm.Event) { server.SelectAction() },
			"reg":           func(_ context.Context, _ *fsm.Event) { server.RegH() },
			"log":           func(_ context.Context, _ *fsm.Event) { server.LogH() },
			"logout":        func(_ context.Context, _ *fsm.Event) { server.LogoutH() },
			"exit":          func(_ context.Context, _ *fsm.Event) { server.ExitH() },
			"deleteAccount": func(_ context.Context, _ *fsm.Event) { server.DeleteAccountH() },
		},
	)

	if err := server.f.Event(context.Background(), "go"); err != nil {
		panic(err)
	}
}
