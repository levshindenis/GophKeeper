package handlers

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/levshindenis/GophKeeper/internal/app/tools"
)

func (mh *MyHandler) GetUpdateTime(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("Cookie")
	login, err := mh.GetDB().GetLogin(cookie.Value)
	if err != nil {
		http.Error(w, "Something bad with GetLogin", http.StatusInternalServerError)
		return
	}

	uTime, err := mh.GetDB().GetUpdateTime(login)
	if err != nil {
		http.Error(w, "Something bad with GetUpdateTime", http.StatusInternalServerError)
		return
	}

	if err = godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv(strings.ToUpper(login) + "_SERVER")

	if _, err = w.Write([]byte(tools.Decrypt(uTime, secretKey))); err != nil {
		http.Error(w, "Something bad with Write", http.StatusInternalServerError)
		return
	}
}
