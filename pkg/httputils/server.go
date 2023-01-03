package httputils

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/saas-templates/go-svelte/pkg/errors"
	"github.com/saas-templates/go-svelte/pkg/log"
)

// Respond writes an HTTP response to the client.
func Respond(wr http.ResponseWriter, req *http.Request, status int, v interface{}) {
	if err, isErr := v.(error); isErr {
		switch {
		case errors.Is(err, errors.ErrInvalid):
			status = http.StatusBadRequest

		case errors.Is(err, errors.ErrNotFound):
			status = http.StatusNotFound

		case errors.Is(err, errors.ErrConflict):
			status = http.StatusConflict

		case errors.Is(err, errors.ErrUnsupported):
			status = http.StatusUnprocessableEntity

		default:
			status = http.StatusInternalServerError
		}

		v = errors.E(err)
	}

	wr.Header().Set("Content-Type", "application/json; charset=utf-8")
	wr.WriteHeader(status)
	_ = json.NewEncoder(wr).Encode(v)
}

// GracefulServe starts HTTP server on addr. Server shuts down gracefully when
// context is cancelled.
func GracefulServe(ctx context.Context, gracePeriod time.Duration, addr string, h http.Handler) error {
	srv := &http.Server{
		Addr:    addr,
		Handler: h,
	}

	go func() {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), gracePeriod)
		defer cancel()

		log.Warnf(ctx, "server shutting down (reason: context_cancelled)")
		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Errorf(ctx, "graceful shutdown failed: %v", err)
			return
		}
		log.Infof(ctx, "graceful shutdown complete")
	}()

	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
