package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/weareinit/Opal/internal/config"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join(string(config.FilesDir), "/dashboard/index.html")
	http.ServeFile(w, r, filePath)
}
