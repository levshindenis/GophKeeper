package main

import (
	"github.com/levshindenis/GophKeeper/internal/app/router"
	"log"
	"net/http"

	"github.com/levshindenis/GophKeeper/internal/app/config"
	"github.com/levshindenis/GophKeeper/internal/app/handlers"
)

func main() {
	var (
		conf config.Config
		h    handlers.MyHandler
	)

	if err := conf.Parse(); err != nil {
		panic(err)
	}

	if conf.GetDBAddress() == "" || conf.GetServerAddress() == "" {
		log.Fatalf("Input config params")
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
