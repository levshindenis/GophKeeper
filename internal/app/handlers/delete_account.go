package handlers

import (
	"net/http"
)

func (mh *MyHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("Cookie")

	login, err := mh.GetDB().GetLogin(cookie.Value)
	if err != nil {
		http.Error(w, "Something bad with GetLogin", http.StatusBadRequest)
		return
	}

	if err := mh.GetDB().DeleteAccount(login, "server"); err != nil {
		http.Error(w, "Something bad with delete data", http.StatusInternalServerError)
		return
	}

	c := &http.Cookie{
		Name:   cookie.Name,
		Value:  "",
		Path:   "/",
		MaxAge: -1}

	http.SetCookie(w, c)

	w.WriteHeader(http.StatusOK)
}
