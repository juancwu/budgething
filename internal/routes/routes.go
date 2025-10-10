package routes

import (
	"io/fs"
	"net/http"

	"git.juancwu.dev/juancwu/budgething/assets"
	"git.juancwu.dev/juancwu/budgething/internal/app"
)

func SetupRoutes(a *app.App) http.Handler {
	mux := http.NewServeMux()

	// Static
	sub, _ := fs.Sub(assets.AssetsFS, ".")
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(sub))))

	return mux
}
