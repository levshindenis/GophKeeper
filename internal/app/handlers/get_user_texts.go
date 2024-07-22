package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/levshindenis/GophKeeper/internal/app/tools"
)

func (mh *MyHandler) GetUserTexts(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("Cookie")
	login, err := mh.GetDB().GetLogin(cookie.Value)
	if err != nil {
		http.Error(w, "Something bad with GetLogin", http.StatusBadRequest)
		return
	}

	allTexts, err := mh.GetDB().GetUserTexts(login)
	if err != nil {
		http.Error(w, "Something bad with GetUserTexts", http.StatusInternalServerError)
		return
	}

	if err = godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv(strings.ToUpper(login) + "_SERVER")

	for i := range allTexts {
		allTexts[i].Name = tools.Decrypt(allTexts[i].Name, secretKey)
		allTexts[i].Description = tools.Decrypt(allTexts[i].Description, secretKey)
		allTexts[i].Comment = tools.Decrypt(allTexts[i].Comment, secretKey)
		allTexts[i].Favourite = tools.Decrypt(allTexts[i].Favourite, secretKey)
	}

	marsh, err := json.Marshal(allTexts)
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
