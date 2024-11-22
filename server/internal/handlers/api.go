package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/weareinit/Opal/internal/auth"
	"github.com/weareinit/Opal/internal/config"
	"github.com/weareinit/Opal/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(middleware.CORSMiddleware)
  r.Use(removeTrailingSlashMiddleware)

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

// we have this func because there are issues with routes that end with and without slashes
// for example: /admin/ would show the admin page when the user shouldn't but /admin would work as intended
func removeTrailingSlashMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" && r.URL.Path[len(r.URL.Path)-1] == '/' {
			newPath := r.URL.Path[:len(r.URL.Path)-1]
			if r.URL.RawQuery != "" {
				newPath = newPath + "?" + r.URL.RawQuery
			}
			http.Redirect(w, r, newPath, http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	})
}
