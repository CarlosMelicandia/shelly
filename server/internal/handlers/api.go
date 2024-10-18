package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/weareinit/Opal/internal/auth"
	"github.com/weareinit/Opal/internal/middleware"
	// "net/http"
)

func Handler(r *chi.Mux) {
	// middleware to strip trailing slashes
	r.Use(chimiddle.StripSlashes)

	// r.Route("/", func(router chi.Router) {
	// 	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 		http.ServeFile(w, r, "../../../client/dist")
	// 	})
	// })

	r.Route("/api/auth", func(router chi.Router) {
		router.Get("/login", oauth.HandleGoogleLogin)
		router.Get("/callback/google", oauth.HandleGoogleCallback)
	})

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
