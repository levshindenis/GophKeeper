package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/levshindenis/GophKeeper/internal/app/models"
	"github.com/levshindenis/GophKeeper/internal/app/tools"
)

func (mh *MyHandler) Register(w http.ResponseWriter, r *http.Request) {
	var (
		dec models.Register
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

	if dec.Login == "" || dec.Password == "" || dec.Word == "" {
		http.Error(w, "Wrong data", http.StatusBadRequest)
		return
	}

	flag, userId, err := mh.GetDB().AddUser(dec)
	if err != nil {
		http.Error(w, "Something bad with AddUser", http.StatusInternalServerError)
		return
	}

	if flag {
		http.Error(w, "Repeated login", http.StatusBadRequest)
		return
	}

	cookie, err := tools.GenerateCookie(userId)
	if err != nil {
		http.Error(w, "Something bad generate cookie", http.StatusInternalServerError)
		return
	}

	//if err = mh.GetCloud().CreateBucket(userId + "-cloud"); err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	http.SetCookie(w, &http.Cookie{Name: "Cookie", Value: cookie})
	mh.GetCookie().Add(cookie, userId)

	w.WriteHeader(http.StatusCreated)
}
