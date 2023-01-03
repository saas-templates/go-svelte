package frontend

import (
	"embed"
	"io/fs"
)

// DistFS holds the static files that need to be served
// as part of the frontend.
var DistFS = mustSubFS(distFS, "dist")

//go:embed dist/**
var distFS embed.FS

func mustSubFS(efs embed.FS, prefix string) fs.FS {
	f, err := fs.Sub(efs, prefix)
	if err != nil {
		panic(err)
	}
	return f
}
