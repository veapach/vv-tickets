package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handlerFunc func(http.ResponseWriter, *http.Request) error

func RegisterRoutes(r *chi.Mux) {

	home := homeHandler{}

	r.Get("/", handler(home.handleIndex))
}

func handler(h handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			handleError(err)
		}
	}
}

func handleError(err error) {
	slog.Error("error during request", slog.String("err", err.Error()))
}
