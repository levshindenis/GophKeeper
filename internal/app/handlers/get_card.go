package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (mh *MyHandler) GetCard(w http.ResponseWriter, r *http.Request) {
	var arr []string

	arr = strings.Split(r.URL.Path, "/")

	cookie, _ := r.Cookie("Cookie")
	userId := mh.GetCookie().GetUserId(cookie.Value)

	item, err := mh.GetDB().GetCard(userId, arr[len(arr)-1])
	if err != nil {
		http.Error(w, "Something bad with GetText", http.StatusInternalServerError)
		return
	}

	resp, err := json.MarshalIndent(item, "", "    ")
	if err != nil {
		http.Error(w, "Something bad with Marshal", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if _, err = w.Write(resp); err != nil {
		http.Error(w, "Something bad with write to ResponseWriter", http.StatusInternalServerError)
		return
	}
}
