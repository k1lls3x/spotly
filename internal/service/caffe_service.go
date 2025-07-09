package service

import (
	"context"
	"spotly/internal/model"
	"spotly/internal/repository"

	"github.com/rs/zerolog"
)

type CaffeService struct {
	repo repository.CaffeRepository
	log  zerolog.Logger
}

func NewCaffeService(repo repository.CaffeRepository, log zerolog.Logger) *CaffeService {
	return &CaffeService{
		repo: repo,
		log:  log,
	}
}

func (s *CaffeService) ListPlaces(ctx context.Context, citySlug, categorySlug string) ([]model.PlaceDTO, error) {
	s.log.Debug().
		Str("city", citySlug).
		Str("category", categorySlug).
		Msg("Fetching places")

	places, err := s.repo.List(ctx, citySlug, categorySlug)
	if err != nil {
		s.log.Error().Err(err).Msg("failed to list places")
		return nil, err
	}
	return places, nil
}

