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

func (mh *MyHandler) AddCards(w http.ResponseWriter, r *http.Request) {
	var (
		dec []models.Card
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
		http.Error(w, "Something bad with add text", http.StatusInternalServerError)
		return
	}

	if err = godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv(strings.ToUpper(login) + "_SERVER")

	for i := range dec {
		dec[i].Bank = tools.Encrypt(dec[i].Bank, secretKey)
		dec[i].Number = tools.Encrypt(dec[i].Number, secretKey)
		dec[i].Month = tools.Encrypt(dec[i].Month, secretKey)
		dec[i].Year = tools.Encrypt(dec[i].Year, secretKey)
		dec[i].CVV = tools.Encrypt(dec[i].CVV, secretKey)
		dec[i].Owner = tools.Encrypt(dec[i].Owner, secretKey)
		dec[i].Comment = tools.Encrypt(dec[i].Comment, secretKey)
		dec[i].Favourite = tools.Encrypt(dec[i].Favourite, secretKey)
	}

	if err = mh.GetDB().AddCards(login, dec); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
