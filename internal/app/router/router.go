package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/levshindenis/GophKeeper/internal/app/handlers"
	"github.com/levshindenis/GophKeeper/internal/app/middleware"
)

func Router(h handlers.MyHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Post("/register", middleware.RegLog(h.Register, &h))
		r.Post("/login", middleware.RegLog(h.Login, &h))
		r.Get("/user/logout", middleware.CheckCookie(h.Logout, &h))
		r.Get("/user/delete-account", middleware.CheckCookie(h.DeleteAccount, &h))
		r.Post("/user/add-texts", middleware.CheckCookie(h.AddTexts, &h))
		r.Post("/user/add-files", middleware.CheckCookie(h.AddFiles, &h))
		r.Post("/user/add-cards", middleware.CheckCookie(h.AddCards, &h))
		r.Get("/user/all-texts", middleware.CheckCookie(h.GetUserTexts, &h))
		r.Get("/user/all-files", middleware.CheckCookie(h.GetUserFiles, &h))
		r.Get("/user/all-cards", middleware.CheckCookie(h.GetUserCards, &h))
		r.Get("/user/clear-data", middleware.CheckCookie(h.ClearUserData, &h))
		r.Get("/user/update-time", middleware.CheckCookie(h.GetUpdateTime, &h))
		r.Post("/user/set-time", middleware.CheckCookie(h.SetUpdateTime, &h))
	})
	return r
}
