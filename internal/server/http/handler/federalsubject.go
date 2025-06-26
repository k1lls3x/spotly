package handler

import (
  "encoding/json"
  "net/http"

  "spotly/internal/model"
  "spotly/internal/service"

  "github.com/go-chi/chi/v5"
  "github.com/google/uuid"
  "github.com/rs/zerolog"
)

type FederalSubjectHandler struct {
  svc *service.FederalSubjectService
  log zerolog.Logger
}

func NewFederalSubjectHandler(svc *service.FederalSubjectService, log zerolog.Logger) *FederalSubjectHandler {
  return &FederalSubjectHandler{svc: svc, log: log}
}

func (h *FederalSubjectHandler) Create(w http.ResponseWriter, r *http.Request) {
  var fs model.FederalSubject
  if err := json.NewDecoder(r.Body).Decode(&fs); err != nil {
    http.Error(w, "invalid payload", http.StatusBadRequest)
    return
  }
  if err := h.svc.Create(r.Context(), &fs); err != nil {
    h.log.Error().Err(err).Msg("Create FederalSubject failed")
    http.Error(w, "internal error", http.StatusInternalServerError)
    return
  }
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(fs)
}

func (h *FederalSubjectHandler) List(w http.ResponseWriter, r *http.Request) {
  list, err := h.svc.List(r.Context())
  if err != nil {
    h.log.Error().Err(err).Msg("List FederalSubject failed")
    http.Error(w, "internal error", http.StatusInternalServerError)
    return
  }
  json.NewEncoder(w).Encode(list)
}

func (h *FederalSubjectHandler) GetByID(w http.ResponseWriter, r *http.Request) {
  id, err := uuid.Parse(chi.URLParam(r, "id"))
  if err != nil {
    http.Error(w, "invalid id", http.StatusBadRequest)
    return
  }
  fs, err := h.svc.GetByID(r.Context(), id)
  if err != nil {
    h.log.Error().Err(err).Msg("GetByID FederalSubject failed")
    http.Error(w, "not found", http.StatusNotFound)
    return
  }
  json.NewEncoder(w).Encode(fs)
}

func (h *FederalSubjectHandler) Update(w http.ResponseWriter, r *http.Request) {
  id, err := uuid.Parse(chi.URLParam(r, "id"))
  if err != nil {
    http.Error(w, "invalid id", http.StatusBadRequest)
    return
  }
  var fs model.FederalSubject
  if err := json.NewDecoder(r.Body).Decode(&fs); err != nil {
    http.Error(w, "invalid payload", http.StatusBadRequest)
    return
  }
  fs.ID = id
  if err := h.svc.Update(r.Context(), &fs); err != nil {
    h.log.Error().Err(err).Msg("Update FederalSubject failed")
    http.Error(w, "internal error", http.StatusInternalServerError)
    return
  }
  json.NewEncoder(w).Encode(fs)
}

func (h *FederalSubjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
  id, err := uuid.Parse(chi.URLParam(r, "id"))
  if err != nil {
    http.Error(w, "invalid id", http.StatusBadRequest)
    return
  }
  if err := h.svc.Delete(r.Context(), id); err != nil {
    h.log.Error().Err(err).Msg("Delete FederalSubject failed")
    http.Error(w, "internal error", http.StatusInternalServerError)
    return
  }
  w.WriteHeader(http.StatusNoContent)
}
