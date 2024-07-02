package main

import (
	"context"
	"fmt"
)

func (s *Server) LogoutH() {
	resp, err := s.client.R().Get(s.address + "/user/logout")
	if err != nil {
		panic(err)
	}

	if resp.StatusCode() == 200 {
		s.cookie = ""
	}

	fmt.Println("Ответ:")

	fmt.Println(resp.Status())
	fmt.Println(resp.String())

	if err = s.f.Event(context.Background(), "mainpage"); err != nil {
		panic(err)
	}
}
