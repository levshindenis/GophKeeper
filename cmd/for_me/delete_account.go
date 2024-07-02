package main

import (
	"context"
	"fmt"
)

func (s *Server) DeleteAccountH() {
	resp, err := s.client.R().Get(s.address + "/user/delete-account")
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
