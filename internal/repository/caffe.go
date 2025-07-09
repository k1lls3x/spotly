package repository

import (
	"context"
	"fmt"
	"spotly/internal/model"

	_"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type CaffeRepository interface {
	List(ctx context.Context, citySlug, categorySlug string) ([]model.PlaceDTO, error)

}

type PgCaffeRepository struct {
	db *sqlx.DB
	log zerolog.Logger
}

func NewCaffeRepository(db *sqlx.DB, log zerolog.Logger) *PgCaffeRepository {
	return &PgCaffeRepository{db: db,
		log : log,
	}
}

func (r *PgCaffeRepository) List(ctx context.Context, citySlug, categorySlug string) ([]model.PlaceDTO, error) {
	query := `
		SELECT
			p.id,
			p.slug,
			(p.i18n->'ru'->>'title')  AS title,
			(p.i18n->'ru'->>'desc')   AS description,
			(p.i18n->'ru'->>'addr')   AS address,
			p.average_check,
			p.rating,
			m.url AS preview_image
		FROM places p
		JOIN place_categories pc ON p.id = pc.place_id
		JOIN categories c ON pc.category_id = c.id
		JOIN cities ci ON p.city_id = ci.id
		LEFT JOIN entity_media em ON em.entity_type = 'place' AND em.entity_id = p.id
		LEFT JOIN media m ON em.media_id = m.id AND m.position = 0
		WHERE p.is_deleted = false
		  AND ci.slug = $1
		  AND c.slug = $2
		ORDER BY p.rating DESC NULLS LAST
	`

	var places []model.PlaceDTO
	if err := r.db.SelectContext(ctx, &places, query, citySlug, categorySlug); err != nil {
		return nil, fmt.Errorf("List cafes: %w", err)
	}
	return places, nil
}
