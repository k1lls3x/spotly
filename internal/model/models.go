package model

import (
    "encoding/json"
    "time"

    "github.com/google/uuid"
)

type JSONB = json.RawMessage

// 1. Субъекты федерации
type FederalSubject struct {
    ID     uuid.UUID `db:"id" json:"id"`
    NameRU string    `db:"name_ru" json:"name_ru"`
    NameEN string    `db:"name_en" json:"name_en"`
}

// 2. Часовые пояса
type Timezone struct {
    ID   int    `db:"id" json:"id"`
    Name string `db:"name" json:"name"`
}

// 3. Города РФ
type City struct {
    ID                uuid.UUID `db:"id" json:"id"`
    FederalSubjectID  uuid.UUID `db:"federal_subject_id" json:"federal_subject_id"`
    TimezoneID        int       `db:"timezone_id" json:"timezone_id"`
    NameRU            string    `db:"name_ru" json:"name_ru"`
    NameEN            string    `db:"name_en" json:"name_en"`
    Slug              string    `db:"slug" json:"slug"`
    OKATO             *string   `db:"okato" json:"okato"`
    Lat               *float64  `db:"lat" json:"lat"`
    Lon               *float64  `db:"lon" json:"lon"`
    I18N              JSONB     `db:"i18n" json:"i18n"`
    IsDeleted         bool      `db:"is_deleted" json:"is_deleted"`
    CreatedAt         time.Time `db:"created_at" json:"created_at"`
    UpdatedAt         time.Time `db:"updated_at" json:"updated_at"`
}

// 4. Категории
type Category struct {
    ID        uuid.UUID  `db:"id" json:"id"`
    Slug      string     `db:"slug" json:"slug"`
    ParentID  *uuid.UUID `db:"parent_id" json:"parent_id"`
    I18N      JSONB      `db:"i18n" json:"i18n"`
    CreatedAt time.Time  `db:"created_at" json:"created_at"`
    UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
}

// 5. События
type Event struct {
    ID           uuid.UUID  `db:"id" json:"id"`
    CityID       uuid.UUID  `db:"city_id" json:"city_id"`
    Slug         string     `db:"slug" json:"slug"`
    KladrID      *string    `db:"kladr_id" json:"kladr_id"`
    I18N         JSONB      `db:"i18n" json:"i18n"`
    StartsAt     time.Time  `db:"starts_at" json:"starts_at"`
    EndsAt       *time.Time `db:"ends_at" json:"ends_at"`
    SearchVector string     `db:"search_vector" json:"search_vector"`
    IsDeleted    bool       `db:"is_deleted" json:"is_deleted"`
    CreatedAt    time.Time  `db:"created_at" json:"created_at"`
    UpdatedAt    time.Time  `db:"updated_at" json:"updated_at"`
}

// 6. Связь “M:N” событие—категория
type EventCategory struct {
    EventID    uuid.UUID `db:"event_id" json:"event_id"`
    CategoryID uuid.UUID `db:"category_id" json:"category_id"`
}

// 7. Медиаконтент
type Media struct {
    ID        uuid.UUID  `db:"id" json:"id"`
    URL       string     `db:"url" json:"url"`
    Caption   *string    `db:"caption" json:"caption"`
    Position  int16      `db:"position" json:"position"`
    CreatedAt time.Time  `db:"created_at" json:"created_at"`
}

type EntityMedia struct {
    EntityType string    `db:"entity_type" json:"entity_type"`
    EntityID   uuid.UUID `db:"entity_id" json:"entity_id"`
    MediaID    uuid.UUID `db:"media_id" json:"media_id"`
}

// 8. Архив старых событий (точно как Event)
type EventArchive struct {
    ID           uuid.UUID  `db:"id" json:"id"`
    CityID       uuid.UUID  `db:"city_id" json:"city_id"`
    Slug         string     `db:"slug" json:"slug"`
    KladrID      *string    `db:"kladr_id" json:"kladr_id"`
    I18N         JSONB      `db:"i18n" json:"i18n"`
    StartsAt     time.Time  `db:"starts_at" json:"starts_at"`
    EndsAt       *time.Time `db:"ends_at" json:"ends_at"`
    SearchVector string     `db:"search_vector" json:"search_vector"`
    IsDeleted    bool       `db:"is_deleted" json:"is_deleted"`
    CreatedAt    time.Time  `db:"created_at" json:"created_at"`
    UpdatedAt    time.Time  `db:"updated_at" json:"updated_at"`
}
