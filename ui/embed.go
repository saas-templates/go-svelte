package ui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:generate yarn
//go:generate yarn build
//go:embed all:dist/*
var distFS embed.FS

// Handler returns a static file serving handler.
func Handler() http.Handler {
	fileSys := http.FS(mustSubFS(distFS, "dist"))
	return http.FileServer(fileSys)
}

func mustSubFS(efs embed.FS, prefix string) fs.FS {
	f, err := fs.Sub(efs, prefix)
	if err != nil {
		panic(err)
	}
	return f
}
