package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/levshindenis/GophKeeper/internal/app/tools"
)

func (mh *MyHandler) SetUpdateTime(w http.ResponseWriter, r *http.Request) {
	var localTime []byte

	localTime, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Something bad with Read Body", http.StatusInternalServerError)
		return
	}

	cookie, _ := r.Cookie("Cookie")
	login, err := mh.GetDB().GetLogin(cookie.Value)
	if err != nil {
		http.Error(w, "Something bad with GetLogin", http.StatusInternalServerError)
		return
	}

	if err = godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv(strings.ToUpper(login) + "_SERVER")

	if err = mh.GetDB().SetUpdateTime(login, tools.Encrypt(string(localTime), secretKey)); err != nil {
		http.Error(w, "Something bad with SetUpdateTime", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
