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
		r.Post("/user/change-texts", middleware.CheckCookie(h.ChangeTexts, &h))
		r.Post("/user/change-files", middleware.CheckCookie(h.ChangeFiles, &h))
		r.Post("/user/change-cards", middleware.CheckCookie(h.ChangeCards, &h))
		r.Post("/user/delete-texts", middleware.CheckCookie(h.DeleteTexts, &h))
		r.Post("/user/delete-files", middleware.CheckCookie(h.DeleteFiles, &h))
		r.Post("/user/delete-cards", middleware.CheckCookie(h.DeleteCards, &h))
		r.Get("/user/list-texts", middleware.CheckCookie(h.ListTexts, &h))
		r.Get("/user/list-files", middleware.CheckCookie(h.ListFiles, &h))
		r.Get("/user/list-cards", middleware.CheckCookie(h.ListCards, &h))
		r.Get("/user/list-favourites", middleware.CheckCookie(h.ListFavourites, &h))
	})
	return r
}
