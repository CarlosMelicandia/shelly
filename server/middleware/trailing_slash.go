package middleware

import (
	"net/http"
)

func RemoveTrailingSlashMiddleware(next http.Handler) http.Handler {
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
