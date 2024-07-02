package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/levshindenis/GophKeeper/internal/app/models"
)

func (mh *MyHandler) ChangeTexts(w http.ResponseWriter, r *http.Request) {
	var (
		dec []models.ChText
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
	userId := mh.GetCookie().GetUserId(cookie.Value)

	if err := mh.GetDB().ChangeTexts(userId, dec); err != nil {
		http.Error(w, "Something bad with change text", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
