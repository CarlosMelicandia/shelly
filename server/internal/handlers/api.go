package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/weareinit/Opal/internal/auth"
	"github.com/weareinit/Opal/internal/config"
	"github.com/weareinit/Opal/internal/middleware"
)

func Handler(r *chi.Mux) {
	r.Route("/", func(router chi.Router) {
		config.FileServer(router, "/", config.FilesDir)
	})

	r.Route("/api/auth", func(router chi.Router) {
		router.Get("/login", oauth.HandleGoogleLogin)
		router.Get("/callback/google", oauth.HandleGoogleCallback)
	})

	r.Route("/dashboard/", func(router chi.Router) {
		router.Use(middleware.JWTMiddleware)
		router.Get("/", DashboardHandler)
	})
}

// todo: move this somewhere else and maybe make the contents an abstraction
func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join(string(config.FilesDir), "/dashboard/index.html")
	http.ServeFile(w, r, filePath)
}
