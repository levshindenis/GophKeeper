package handlers

import "net/http"

func (mh *MyHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("Cookie")

	userId := mh.GetCookie().GetUserId(cookie.Value)

	if err := mh.GetCloud().DeleteBucket(userId + "ooo"); err != nil {
		http.Error(w, "Something bad with delete bucket", http.StatusInternalServerError)
		return
	}

	if err := mh.GetDB().DeleteAccount(userId); err != nil {
		http.Error(w, "Something bad with delete data", http.StatusInternalServerError)
		return
	}

	mh.GetCookie().Delete(cookie.Value)

	w.WriteHeader(http.StatusOK)
}
