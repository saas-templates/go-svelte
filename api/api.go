package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Router(cfg Config) http.Handler {
	r := chi.NewRouter()
	// TODO: add apis here.
	return r
}

type Config struct {
	// TODO: add API related configurations.
}
