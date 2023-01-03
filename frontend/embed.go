package frontend

import (
	"embed"
	"io/fs"
)

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
