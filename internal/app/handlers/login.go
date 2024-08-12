package handlers

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

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
		http.Error(w, "Empty data", http.StatusBadRequest)
		return
	}

	if !mh.GetDB().CheckUser(dec) {
		http.Error(w, "Wrong login or password", http.StatusBadRequest)
		return
	}

	cookie, err := tools.GenerateCookie(strconv.Itoa(rand.Intn(100)))
	if err != nil {
		http.Error(w, "Something bad with generate cookie", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "Cookie", Value: cookie})

	if err = mh.GetDB().SetCookie(cookie, dec.Login); err != nil {
		http.Error(w, "Something bad with SetCookie", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
