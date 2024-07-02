package main

import (
	"fmt"
	"os"
)

func (s *Server) ExitH() {
	_, err := s.client.R().Get(s.address + "/user/logout")
	if err != nil {
		panic(err)
	}

	fmt.Println("Exit")
	os.Exit(0)
}
