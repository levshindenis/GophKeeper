package handlers

import (
	"net/http"
)

func (mh *MyHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("Cookie")

	mh.GetCookie().Delete(cookie.Value)

	c := &http.Cookie{
		Name:   cookie.Name,
		Value:  "",
		Path:   "/",
		MaxAge: -1}

	http.SetCookie(w, c)

	w.WriteHeader(http.StatusOK)
}
