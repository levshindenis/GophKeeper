package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (mh *MyHandler) DeleteFiles(w http.ResponseWriter, r *http.Request) {
	var (
		dec []string
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

	if err := mh.GetCloud().DeleteFiles(userId+"-cloud", dec); err != nil {
		http.Error(w, "Something bad with delete s3 file", http.StatusInternalServerError)
		return
	}

	if err := mh.GetDB().DeleteFiles(userId, dec); err != nil {
		http.Error(w, "Something bad with delete file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
