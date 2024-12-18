// This is how we can fetch specific information depending on the route that we are in.

// For example, if we create a new page on the client (called /fakeRoute), we would need
// to build the client app first using `npm run build` and then create a handler function
// to serve it.

// Create a new page on the client -> Build the client -> Create a route in this file -> Create a handler function to serve the dist file.
// As of writing this (12/15/2024), you can reference how to do this by looking at how the admin or dashboard route was done and its handler function.

package handlers

import (
	"github.com/go-chi/chi"
	"github.com/weareinit/Opal/internal/auth"
	"github.com/weareinit/Opal/internal/config"
	"github.com/weareinit/Opal/middleware"
)

func Handler(r *chi.Mux) {
	// Before every route, we will run these middlewares.
	r.Use(middleware.CORSMiddleware)
	r.Use(middleware.RemoveTrailingSlashMiddleware)

	r.Route("/", func(router chi.Router) {
		config.FileServer(router, "/", config.FilesDir)
	})

	// You probably don't need to touch this route
	r.Route("/api/auth", func(router chi.Router) {
		router.Get("/login/google", oauth.HandleGoogleLogin)
		router.Get("/callback/google", oauth.HandleGoogleCallback)

		router.Get("/login/discord", oauth.HandleDiscordLogin)
		router.Get("/callback/discord", oauth.HandleDiscordCallback)
	})

	// The convention for these routes is to add a handler function. This makes
	// it cleaner and you know exactly what each route does pretty quickly.
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
