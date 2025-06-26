package http

import (
  "net/http"

  "spotly/internal/service"
 "spotly/internal/server/http/handler"
  "github.com/go-chi/chi/v5"
  "github.com/go-chi/chi/v5/middleware"
  "github.com/rs/zerolog"
)

func NewRouter(log zerolog.Logger, fsSvc *service.FederalSubjectService /*, citySvc, … */) http.Handler {
  r := chi.NewRouter()
  r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)

  // Корень
  r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Spotly API"))
  })

  // FederalSubject
  fsHandler := handler.NewFederalSubjectHandler(fsSvc, log)
  r.Route("/federal_subjects", func(r chi.Router) {
    r.Post("/", fsHandler.Create)
    r.Get("/", fsHandler.List)
    r.Get("/{id}", fsHandler.GetByID)
    r.Put("/{id}", fsHandler.Update)
    r.Delete("/{id}", fsHandler.Delete)
  })

  // TODO: маршруты для /cities, /categories, /events, /media
  return r
}
