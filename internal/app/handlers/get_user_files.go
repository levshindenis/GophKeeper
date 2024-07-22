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

func (mh *MyHandler) GetUserFiles(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("Cookie")
	login, err := mh.GetDB().GetLogin(cookie.Value)
	if err != nil {
		http.Error(w, "Something bad with GetLogin", http.StatusBadRequest)
		return
	}

	allFiles, err := mh.GetDB().GetUserFiles(login)
	if err != nil {
		http.Error(w, "Something bad with GetUserTexts", http.StatusInternalServerError)
		return
	}

	if err = godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv(strings.ToUpper(login) + "_SERVER")

	for i := range allFiles {
		allFiles[i].Name = tools.Decrypt(allFiles[i].Name, secretKey)
		allFiles[i].Comment = tools.Decrypt(allFiles[i].Comment, secretKey)
		allFiles[i].Favourite = tools.Decrypt(allFiles[i].Favourite, secretKey)
	}

	marsh, err := json.Marshal(allFiles)
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
