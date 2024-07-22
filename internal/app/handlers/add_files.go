package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/levshindenis/GophKeeper/internal/app/models"
	"github.com/levshindenis/GophKeeper/internal/app/tools"
	"log"
	"net/http"
	"os"
	"strings"
)

func (mh *MyHandler) AddFiles(w http.ResponseWriter, r *http.Request) {
	var (
		dec []models.File
		buf bytes.Buffer
	)

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Wrong data type", http.StatusBadRequest)
		return
	}

	if _, err := buf.ReadFrom(r.Body); err != nil {
		http.Error(w, "Something bad with read body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(buf.Bytes(), &dec); err != nil {
		http.Error(w, "Something bad with decoding JSON", http.StatusInternalServerError)
		return
	}

	cookie, _ := r.Cookie("Cookie")
	login, err := mh.GetDB().GetLogin(cookie.Value)
	if err != nil {
		http.Error(w, "Something bad with GetLogin", http.StatusBadRequest)
		return
	}

	if err = godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv(strings.ToUpper(login) + "_SERVER")

	for i := range dec {
		dec[i].Name = tools.Encrypt(dec[i].Name, secretKey)
		dec[i].Comment = tools.Encrypt(dec[i].Comment, secretKey)
		dec[i].Favourite = tools.Encrypt(dec[i].Favourite, secretKey)
	}

	if err = mh.GetDB().AddFiles(login, dec); err != nil {
		http.Error(w, "Something bad with add text", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
