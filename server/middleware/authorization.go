package middleware

import (
	"context"
	"net/http"

	"github.com/weareinit/Opal/internal/helpers/tokens"
	"github.com/weareinit/Opal/internal/helpers/users"
)

// check if they are currently logged in
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken, err := tokens.ValidateAccessToken(w, r)
		if err != nil {
			// Attempt to refresh the access token using the refresh token
			newAccessToken, err := tokens.RefreshTokens(w, r)
			if err != nil {
				http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "access_token", newAccessToken)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		ctx := context.WithValue(r.Context(), "access_token", accessToken)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AdminMiddleware(next http.Handler) http.Handler {
  return JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    userId, err := users.GetUserId(w, r)
    if err != nil {
        http.Error(w, "Unauthorized: "+err.Error(), http.StatusNotFound)
        return
    }

    isAdmin := users.IsUserAdmin(userId, w, r)
    if isAdmin {
      next.ServeHTTP(w, r)
    } else {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
  }))
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from your frontend origin
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4321")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
