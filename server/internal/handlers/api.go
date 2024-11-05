package handlers

import (
	"github.com/go-chi/chi"
	"github.com/weareinit/Opal/internal/auth"
	"github.com/weareinit/Opal/internal/config"
	"github.com/weareinit/Opal/internal/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(middleware.CORSMiddleware)

	r.Route("/", func(router chi.Router) {
		config.FileServer(router, "/", config.FilesDir)
	})

	r.Route("/api/auth", func(router chi.Router) {
		router.Get("/login", oauth.HandleGoogleLogin)
		router.Get("/callback/google", oauth.HandleGoogleCallback)
	})

	r.Route("/api/user", func(router chi.Router) {
		router.Get("/", UserHandler)
	})

	r.Route("/dashboard/", func(router chi.Router) {
		router.Use(middleware.JWTMiddleware)
		router.Get("/", DashboardHandler)
	})
}
