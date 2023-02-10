package handlers

import (
	"sync/atomic"

	"github.com/go-chi/chi/v5"
)

// Router register necessary routes and returns an instance of a router.
func Router() *chi.Mux {

	r := chi.NewRouter()
	isReady := &atomic.Value{}
	isReady.Store(true)

	r.Post("/", home)
	r.HandleFunc("/healthz", healthz)
	r.HandleFunc("/readyz", readyz(isReady))
	return r
}
