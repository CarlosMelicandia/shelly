package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/weareinit/Opal/internal/middleware"
	"net/http"
)

func Handler(r *chi.Mux) {
	// get rid of slashes in the end of api fetches
	r.Use(chimiddle.StripSlashes)

	r.Route("/", func(router chi.Router) {
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("This is a test to show how its written and it works. We can delete this later"))
		})
	})

	r.Route("/account", func(router chi.Router) {

		// Middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
