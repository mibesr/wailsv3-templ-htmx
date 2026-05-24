package web

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed all:static
var staticFS embed.FS

// StaticHandler serves files from web/static at /static/* (paths relative to web/static).
func StaticHandler() http.Handler {
	sub, err := fs.Sub(staticFS, "static")
	if err != nil {
		panic(err)
	}
	return http.StripPrefix("/static/", http.FileServer(http.FS(sub)))
}
