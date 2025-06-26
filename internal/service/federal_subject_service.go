package service

import (
	"context"
	"spotly/internal/model"
	"spotly/internal/repository"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type FederalSubjectService struct {
	repo repository.FederalSubjectRepository
	log  zerolog.Logger
}

func NewFederalSubjectService(repo repository.FederalSubjectRepository, log zerolog.Logger) *FederalSubjectService {
	return &FederalSubjectService{
		repo: repo,
		log:  log,
	}
}

func (s *FederalSubjectService) Create(ctx context.Context, fs *model.FederalSubject) error {
	// generate UUID if not set
	if fs.ID == uuid.Nil {
		fs.ID = uuid.New()
	}
	if err := s.repo.Create(ctx, fs); err != nil {
		s.log.Error().
			Err(err).
			Str("id", fs.ID.String()).
			Msg("Не удалось создать субъект федерации")
		return err
	}
	s.log.Info().
		Str("id", fs.ID.String()).
		Msg("Субъект федерации успешно создан")
	return nil
}

func (s *FederalSubjectService) GetByID(ctx context.Context, id uuid.UUID) (*model.FederalSubject, error) {
	fs, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.log.Error().
			Err(err).
			Str("id", id.String()).
			Msg("Не удалось получить субъект федерации")
		return nil, err
	}
	s.log.Info().
		Str("id", id.String()).
		Msg("Субъект федерации получен")
	return fs, nil
}

func (s *FederalSubjectService) List(ctx context.Context) ([]model.FederalSubject, error) {
	list, err := s.repo.List(ctx)
	if err != nil {
		s.log.Error().
			Err(err).
			Msg("Не удалось получить список субъектов федерации")
		return nil, err
	}
	s.log.Info().
		Int("count", len(list)).
		Msg("Список субъектов федерации получен")
	return list, nil
}

func (s *FederalSubjectService) Update(ctx context.Context, fs *model.FederalSubject) error {
	if err := s.repo.Update(ctx, fs); err != nil {
		s.log.Error().
			Err(err).
			Str("id", fs.ID.String()).
			Msg("Не удалось обновить субъект федерации")
		return err
	}
	s.log.Info().
		Str("id", fs.ID.String()).
		Msg("Субъект федерации успешно обновлён")
	return nil
}

func (s *FederalSubjectService) Delete(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		s.log.Error().
			Err(err).
			Str("id", id.String()).
			Msg("Не удалось удалить субъект федерации")
		return err
	}
	s.log.Info().
		Str("id", id.String()).
		Msg("Субъект федерации успешно удалён")
	return nil
}
