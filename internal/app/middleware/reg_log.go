package middleware

import (
	"net/http"

	"github.com/levshindenis/GophKeeper/internal/app/handlers"
)

func RegLog(next http.HandlerFunc, hm *handlers.MyHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Cookie")
		if err == nil {
			if hm.GetDB().CheckCookie(cookie.Value) {
				http.Error(w, "You are already logged in", http.StatusBadRequest)
				return
			}
			http.Error(w, "Fake cookie", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}
