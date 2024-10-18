package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/weareinit/Opal/internal/handlers"
)

func main() {
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "../../../client/dist"))

	r.Route("/", func(router chi.Router) {
		FileServer(router, "/", filesDir)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(workDir, "../../../client/dist/index.html"))
	})

	fmt.Println("Starting GO API service on localhost:8000...")

	fmt.Println(`
  ██████  ██░ ██ ▓█████  ██▓     ██▓     ██░ ██  ▄▄▄       ▄████▄   ██ ▄█▀  ██████ 
▒██    ▒ ▓██░ ██▒▓█   ▀ ▓██▒    ▓██▒    ▓██░ ██▒▒████▄    ▒██▀ ▀█   ██▄█▒ ▒██    ▒ 
░ ▓██▄   ▒██▀▀██░▒███   ▒██░    ▒██░    ▒██▀▀██░▒██  ▀█▄  ▒▓█    ▄ ▓███▄░ ░ ▓██▄   
  ▒   ██▒░▓█ ░██ ▒▓█  ▄ ▒██░    ▒██░    ░▓█ ░██ ░██▄▄▄▄██ ▒▓▓▄ ▄██▒▓██ █▄   ▒   ██▒
▒██████▒▒░▓█▒░██▓░▒████▒░██████▒░██████▒░▓█▒░██▓ ▓█   ▓██▒▒ ▓███▀ ░▒██▒ █▄▒██████▒▒
▒ ▒▓▒ ▒ ░ ▒ ░░▒░▒░░ ▒░ ░░ ▒░▓  ░░ ▒░▓  ░ ▒ ░░▒░▒ ▒▒   ▓▒█░░ ░▒ ▒  ░▒ ▒▒ ▓▒▒ ▒▓▒ ▒ ░
░ ░▒  ░ ░ ▒ ░▒░ ░ ░ ░  ░░ ░ ▒  ░░ ░ ▒  ░ ▒ ░▒░ ░  ▒   ▒▒ ░  ░  ▒   ░ ░▒ ▒░░ ░▒  ░ ░
░  ░  ░   ░  ░░ ░   ░     ░ ░     ░ ░    ░  ░░ ░  ░   ▒   ░        ░ ░░ ░ ░  ░  ░  
      ░   ░  ░  ░   ░  ░    ░  ░    ░  ░ ░  ░  ░      ░  ░░ ░      ░  ░         ░  
                                                          ░                        `)
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}

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
