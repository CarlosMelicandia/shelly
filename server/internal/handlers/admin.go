package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/weareinit/Opal/internal/config"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join(string(config.FilesDir), "/admin/index.html")
	http.ServeFile(w, r, filePath)
}

