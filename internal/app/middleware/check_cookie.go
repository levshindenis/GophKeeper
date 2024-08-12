package middleware

import (
	"net/http"

	"github.com/levshindenis/GophKeeper/internal/app/handlers"
)

func CheckCookie(next http.HandlerFunc, hm *handlers.MyHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Cookie")
		if err != nil || !hm.GetDB().CheckCookie(cookie.Value) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
