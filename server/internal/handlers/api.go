package handlers

import (
	"github.com/go-chi/chi"
	"github.com/weareinit/Opal/internal/auth"
	"github.com/weareinit/Opal/internal/config"
	"github.com/weareinit/Opal/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(middleware.CORSMiddleware)
	r.Use(middleware.RemoveTrailingSlashMiddleware)

	r.Route("/", func(router chi.Router) {
		config.FileServer(router, "/", config.FilesDir)
	})

	r.Route("/api/auth", func(router chi.Router) {
		router.Get("/login/google", oauth.HandleGoogleLogin)
		router.Get("/callback/google", oauth.HandleGoogleCallback)

		router.Get("/login/discord", oauth.HandleDiscordLogin)
		router.Get("/callback/discord", oauth.HandleDiscordCallback)
	})

	r.Route("/api/getUser", func(router chi.Router) {
		router.Get("/", GetUserHandler)
	})

	r.Route("/api/getHacker", func(router chi.Router) {
		router.Get("/", GetHackerHandler)
	})

	r.Route("/api/createHacker", func(router chi.Router) {
		router.Post("/", CreateHackerHandler)
	})

	r.Route("/dashboard", func(router chi.Router) {
		router.Get("/", DashboardHandler)
	})

	r.Route("/admin", func(router chi.Router) {
		router.Get("/", AdminHandler)
	})
}
