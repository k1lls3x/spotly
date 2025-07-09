package handler

import (
	"encoding/json"
	"net/http"

	"spotly/internal/service"

	"github.com/rs/zerolog"
)

type PlaceHandler struct {
	svc *service.CaffeService
	log zerolog.Logger
}

func NewPlaceHandler(svc *service.CaffeService, log zerolog.Logger) *PlaceHandler {
	return &PlaceHandler{svc: svc, log: log}
}

func (h *PlaceHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	city := r.URL.Query().Get("city")
	category := r.URL.Query().Get("category")
	if city == "" || category == "" {
		http.Error(w, "missing city or category query param", http.StatusBadRequest)
		return
	}

	places, err := h.svc.ListPlaces(ctx, city, category)
	if err != nil {
		h.log.Error().Err(err).Msg("failed to list places")
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(places)
}
