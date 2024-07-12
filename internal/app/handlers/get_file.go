package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (mh *MyHandler) GetFile(w http.ResponseWriter, r *http.Request) {
	var (
		dec string
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

	item, err := mh.GetDB().GetFile(userId, dec)
	if err != nil {
		http.Error(w, "Something bad with GetFile", http.StatusInternalServerError)
		return
	}

	resp, err := json.MarshalIndent(item, "", "    ")
	if err != nil {
		http.Error(w, "Something bad with Marshal", http.StatusInternalServerError)
		return
	}

	object, err := mh.GetCloud().GetFile(userId, dec)
	if err != nil {
		http.Error(w, "Something bad with Cloud(GetFile)", http.StatusInternalServerError)
		return
	}

	defer object.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err = w.Write(resp); err != nil {
		http.Error(w, "Something bad with write to ResponseWriter", http.StatusInternalServerError)
		return
	}
}
