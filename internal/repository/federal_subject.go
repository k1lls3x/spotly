package repository

import (
	"context"
	"spotly/internal/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

// -----------------------------
// FederalSubjectRepository
// -----------------------------

type FederalSubjectRepository interface {
	Create(ctx context.Context, fs *model.FederalSubject) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.FederalSubject, error)
	List(ctx context.Context) ([]model.FederalSubject, error)
	Update(ctx context.Context, fs *model.FederalSubject) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type PgFederalSubjectRepository struct {
	db  *sqlx.DB
	log zerolog.Logger
}

func NewFederalSubjectRepository(db *sqlx.DB, log zerolog.Logger) *PgFederalSubjectRepository {
	return &PgFederalSubjectRepository{db: db, log: log}
}

func (r *PgFederalSubjectRepository) Create(ctx context.Context, fs *model.FederalSubject) error {
	if fs.ID == uuid.Nil {
		fs.ID = uuid.New()
	}
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO federal_subjects (id, name_ru, name_en)
		VALUES ($1, $2, $3)
	`, fs.ID, fs.NameRU, fs.NameEN)
	if err != nil {
		r.log.Error().Err(err).Msg("Ошибка при создании субъекта федерации")
		return err
	}
	r.log.Info().Str("id", fs.ID.String()).Msg("Субъект федерации создан")
	return nil
}

func (r *PgFederalSubjectRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.FederalSubject, error) {
	var fs model.FederalSubject
	err := r.db.GetContext(ctx, &fs, `
		SELECT id, name_ru, name_en
		FROM federal_subjects
		WHERE id = $1
	`, id)
	if err != nil {
		r.log.Error().Err(err).Msg("Ошибка при получении субъекта федерации")
		return nil, err
	}
	return &fs, nil
}

func (r *PgFederalSubjectRepository) List(ctx context.Context) ([]model.FederalSubject, error) {
	var list []model.FederalSubject
	err := r.db.SelectContext(ctx, &list, `
		SELECT id, name_ru, name_en
		FROM federal_subjects
		ORDER BY name_ru
	`)
	if err != nil {
		r.log.Error().Err(err).Msg("Ошибка при получении списка субъектов федерации")
		return nil, err
	}
	return list, nil
}

func (r *PgFederalSubjectRepository) Update(ctx context.Context, fs *model.FederalSubject) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE federal_subjects
		SET name_ru = $2, name_en = $3
		WHERE id = $1
	`, fs.ID, fs.NameRU, fs.NameEN)
	if err != nil {
		r.log.Error().Err(err).Msg("Ошибка при обновлении субъекта федерации")
		return err
	}
	r.log.Info().Str("id", fs.ID.String()).Msg("Субъект федерации обновлён")
	return nil
}

func (r *PgFederalSubjectRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, `
		DELETE FROM federal_subjects
		WHERE id = $1
	`, id)
	if err != nil {
		r.log.Error().Err(err).Msg("Ошибка при удалении субъекта федерации")
		return err
	}
	r.log.Info().Str("id", id.String()).Msg("Субъект федерации удалён")
	return nil
}
