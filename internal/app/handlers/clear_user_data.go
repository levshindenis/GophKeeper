package handlers

import (
	"net/http"
)

func (mh *MyHandler) ClearUserData(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("Cookie")
	login, err := mh.GetDB().GetLogin(cookie.Value)
	if err != nil {
		http.Error(w, "Something bad with GetLogin", http.StatusBadRequest)
		return
	}

	if err = mh.GetDB().DeleteTexts(login, nil, "all"); err != nil {
		http.Error(w, "Something bad with DeleteTexts", http.StatusInternalServerError)
		return
	}
	if err = mh.GetDB().DeleteFiles(login, nil, "all"); err != nil {
		http.Error(w, "Something bad with DeleteFiles", http.StatusInternalServerError)
		return
	}
	if err = mh.GetDB().DeleteCards(login, nil, "all"); err != nil {
		http.Error(w, "Something bad with DeleteCards", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
