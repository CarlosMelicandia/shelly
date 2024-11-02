package config

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
)

var (
	WorkDir, _ = os.Getwd()
	// FilesDir points to the client distribution folder
	FilesDir = http.Dir(filepath.Join(WorkDir, "../../../client/dist"))
)

// FileServer serves static files from a given directory
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	fs := http.StripPrefix(path, http.FileServer(root))
	r.Get(path+"*", func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
}
