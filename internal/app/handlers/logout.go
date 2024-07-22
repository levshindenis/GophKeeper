package handlers

import (
	"net/http"
)

func (mh *MyHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("Cookie")

	c := &http.Cookie{
		Name:   cookie.Name,
		Value:  "",
		Path:   "/",
		MaxAge: -1}

	http.SetCookie(w, c)

	if err := mh.GetDB().SetCookie(cookie.Value, ""); err != nil {
		http.Error(w, "Something bas with delete cookie from db", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
