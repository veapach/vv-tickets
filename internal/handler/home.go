package handler

import (
	"net/http"

	"github.com/veapach/vv-tickets/internal/view/home"
)

type homeHandler struct{}

func (h homeHandler) handleIndex(w http.ResponseWriter, r *http.Request) error {
	return home.Index().Render(r.Context(), w)
}

func (h homeHandler) handleAbout(w http.ResponseWriter, r *http.Request) error {
	return home.About().Render(r.Context(), w)
}
