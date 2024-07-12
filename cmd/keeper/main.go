package main

import (
	"log"
	"net/http"

	"github.com/levshindenis/GophKeeper/internal/app/config"
	"github.com/levshindenis/GophKeeper/internal/app/handlers"
	"github.com/levshindenis/GophKeeper/internal/app/router"
)

func main() {
	var (
		conf config.Config
		h    handlers.MyHandler
	)

	if err := conf.ParseFlags(); err != nil {
		panic(err)
	}

	if conf.GetDBAddress() == "" {
		log.Fatalf("Input db address")
	}

	if err := h.Init(conf); err != nil {
		panic(err)
	}

	log.Println("Server start")

	if err := http.ListenAndServe(conf.GetServerAddress(), router.Router(h)); err != nil {
		panic(err)
	}

	if err := h.Cancel(); err != nil {
		panic(err)
	}
}
