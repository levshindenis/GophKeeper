package main

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *Server) ListCardsH() {
	var arr []string

	resp, err := s.client.R().Get(s.address + "/user/list-cards")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(resp.Body(), &arr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Cards:")
	for i := range arr {
		fmt.Println(arr[i])
	}

	if err = s.f.Event(context.Background(), "mainpage"); err != nil {
		panic(err)
	}
}
