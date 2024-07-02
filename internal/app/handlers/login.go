package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/levshindenis/GophKeeper/internal/app/models"
	"github.com/levshindenis/GophKeeper/internal/app/tools"
)

func (mh *MyHandler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		dec models.Login
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

	if dec.Login == "" || dec.Password == "" {
		http.Error(w, "Wrong data", http.StatusBadRequest)
		return
	}

	userId, err := mh.GetDB().CheckUser(dec)
	if err != nil {
		http.Error(w, "Wrong login or password", http.StatusBadRequest)
		return
	}

	cookie, err := tools.GenerateCookie(userId)
	if err != nil {
		http.Error(w, "Something bad generate cookie", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "Cookie", Value: cookie})
	mh.GetCookie().Add(cookie, userId)

	w.WriteHeader(http.StatusOK)
}
