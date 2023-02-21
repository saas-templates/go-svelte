package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type Config struct {
	// TODO: add API related configurations.
}

func Router(cfg Config) http.Handler {
	r := chi.NewRouter()

	r.Get("/data", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		c, err := r.Cookie("foo")
		if err != nil || c == nil {
			http.SetCookie(w, &http.Cookie{
				Name:  "foo",
				Value: "This is not your first time!",
			})

			_ = json.NewEncoder(w).Encode(map[string]any{
				"message": "Welcome firs time user!",
			})
		} else {
			_ = json.NewEncoder(w).Encode(map[string]any{
				"message": c.Value,
			})
		}
	})

	return r
}
