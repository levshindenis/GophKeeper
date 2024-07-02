package handlers

import (
	"encoding/json"
	"net/http"
)

func (mh *MyHandler) ListFavourites(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("Cookie")
	userId := mh.GetCookie().GetUserId(cookie.Value)

	items, err := mh.GetDB().ListFavourites(userId)
	if err != nil {
		http.Error(w, "Something bad with ListFavourites", http.StatusInternalServerError)
		return
	}

	resp, err := json.MarshalIndent(items, "", "    ")
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
