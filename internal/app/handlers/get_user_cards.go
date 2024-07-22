package handlers

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/levshindenis/GophKeeper/internal/app/tools"
	"log"
	"net/http"
	"os"
	"strings"
)

func (mh *MyHandler) GetUserCards(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("Cookie")
	login, err := mh.GetDB().GetLogin(cookie.Value)
	if err != nil {
		http.Error(w, "Something bad with GetLogin", http.StatusBadRequest)
		return
	}

	allCards, err := mh.GetDB().GetUserCards(login)
	if err != nil {
		http.Error(w, "Something bad with GetUserTexts", http.StatusInternalServerError)
		return
	}

	if err = godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv(strings.ToUpper(login) + "_SERVER")

	for i := range allCards {
		allCards[i].Bank = tools.Decrypt(allCards[i].Bank, secretKey)
		allCards[i].Number = tools.Decrypt(allCards[i].Number, secretKey)
		allCards[i].Month = tools.Decrypt(allCards[i].Month, secretKey)
		allCards[i].Year = tools.Decrypt(allCards[i].Year, secretKey)
		allCards[i].CVV = tools.Decrypt(allCards[i].CVV, secretKey)
		allCards[i].Owner = tools.Decrypt(allCards[i].Owner, secretKey)
		allCards[i].Comment = tools.Decrypt(allCards[i].Comment, secretKey)
		allCards[i].Favourite = tools.Decrypt(allCards[i].Favourite, secretKey)
	}

	marsh, err := json.Marshal(allCards)
	if err != nil {
		http.Error(w, "Something bad with Marshal", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	if _, err = w.Write(marsh); err != nil {
		http.Error(w, "Something bad with Write", http.StatusInternalServerError)
		return
	}
}
