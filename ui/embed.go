package ui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed all:dist/**
var distFS embed.FS

func Handler() http.Handler {
	f, err := fs.Sub(distFS, "dist")
	if err != nil {
		panic(err)
	}
	return http.FileServer(http.FS(f))
}
