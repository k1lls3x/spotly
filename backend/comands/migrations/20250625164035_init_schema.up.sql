-- =======================================
--  Итоговая схема «Только Россия»
--  PostgreSQL ≥ 15
--  Поддержка UUID, мультиязычности, аудита, медиа
-- =======================================

-- Расширения
CREATE EXTENSION IF NOT EXISTS pgcrypto;       -- для gen_random_uuid()
CREATE EXTENSION IF NOT EXISTS pg_trgm;         -- для продвинутого поиска
CREATE EXTENSION IF NOT EXISTS pg_cron;         -- для планировщика фоновых задач

-- 1. Субъекты федерации
CREATE TABLE federal_subjects (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name_ru     TEXT NOT NULL UNIQUE,
    name_en     TEXT NOT NULL
);

-- 2. Часовые пояса
CREATE TABLE timezones (
    id      SERIAL PRIMARY KEY,
    name    TEXT NOT NULL UNIQUE  -- e.g. 'Europe/Moscow'
);

-- 3. Города РФ
CREATE TABLE cities (
    id                   UUID    PRIMARY KEY DEFAULT gen_random_uuid(),
    federal_subject_id   UUID    NOT NULL REFERENCES federal_subjects(id) ON DELETE RESTRICT,
    timezone_id          INT     NOT NULL REFERENCES timezones(id) ON DELETE RESTRICT,
    name_ru              TEXT    NOT NULL,
    name_en              TEXT    NOT NULL,
    slug                 TEXT    NOT NULL UNIQUE,
    okato                TEXT,
    lat                  NUMERIC(9,6) CHECK (lat BETWEEN -90 AND 90),
    lon                  NUMERIC(9,6) CHECK (lon BETWEEN -180 AND 180),
    i18n                 JSONB   DEFAULT '{}',
    is_deleted           BOOLEAN DEFAULT FALSE,
    created_at           TIMESTAMPTZ DEFAULT now(),
    updated_at           TIMESTAMPTZ DEFAULT now(),
    UNIQUE (name_ru, federal_subject_id)
);

-- 4. Категории с иерархией и мультиязычностью
CREATE TABLE categories (
    id          UUID    PRIMARY KEY DEFAULT gen_random_uuid(),
    slug        TEXT    NOT NULL UNIQUE,
    parent_id   UUID    REFERENCES categories(id) ON DELETE SET NULL,
    i18n        JSONB   DEFAULT '{}',   -- {"ru": {"title":..., "desc":...}}
    created_at  TIMESTAMPTZ DEFAULT now(),
    updated_at  TIMESTAMPTZ DEFAULT now()
);

-- 5. События
CREATE TABLE events (
    id                  UUID    PRIMARY KEY DEFAULT gen_random_uuid(),
    city_id             UUID    NOT NULL REFERENCES cities(id) ON DELETE CASCADE,
    slug                TEXT    NOT NULL UNIQUE,
    kladr_id            TEXT,
    i18n                JSONB   DEFAULT '{}',  -- {"ru": {"title":..., "desc":..., "addr":...}}
    starts_at           TIMESTAMPTZ NOT NULL,
    ends_at             TIMESTAMPTZ,
    search_vector       tsvector,
    is_deleted          BOOLEAN DEFAULT FALSE,
    created_at          TIMESTAMPTZ DEFAULT now(),
    updated_at          TIMESTAMPTZ DEFAULT now()
);

-- Ограничение на даты
ALTER TABLE events
  ADD CONSTRAINT chk_event_dates CHECK (ends_at IS NULL OR ends_at >= starts_at);

-- Индексы
CREATE INDEX idx_events_city_date     ON events(city_id, starts_at);
CREATE INDEX idx_events_starts        ON events(starts_at);
CREATE INDEX idx_events_search        ON events USING GIN (search_vector);

-- 6. Связь «M:N» событие—категория
CREATE TABLE event_categories (
    event_id     UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    category_id  UUID NOT NULL REFERENCES categories(id) ON DELETE RESTRICT,
    PRIMARY KEY (event_id, category_id)
);
CREATE INDEX idx_event_categories_cat   ON event_categories(category_id);
CREATE INDEX idx_event_categories_evt   ON event_categories(event_id);

-- 7. Медиаконтент (универсальный)
CREATE TABLE media (
    id          UUID    PRIMARY KEY DEFAULT gen_random_uuid(),
    url         TEXT    NOT NULL,
    caption     TEXT,
    position    SMALLINT DEFAULT 0,
    created_at  TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE entity_media (
    entity_type TEXT    NOT NULL,
    entity_id   UUID    NOT NULL,
    media_id    UUID    NOT NULL REFERENCES media(id) ON DELETE CASCADE,
    PRIMARY KEY (entity_type, entity_id, media_id)
);

-- 8. Архивирование старых событий
CREATE TABLE events_archive (LIKE events INCLUDING ALL);

-- 9. Триггеры: обновление updated_at
CREATE FUNCTION fn_set_updated_at() RETURNS trigger LANGUAGE plpgsql AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END; $$;
CREATE TRIGGER trg_cities_updated_at    BEFORE UPDATE ON cities    FOR EACH ROW EXECUTE FUNCTION fn_set_updated_at();
CREATE TRIGGER trg_categories_updated_at BEFORE UPDATE ON categories FOR EACH ROW EXECUTE FUNCTION fn_set_updated_at();
CREATE TRIGGER trg_events_updated_at     BEFORE UPDATE ON events     FOR EACH ROW EXECUTE FUNCTION fn_set_updated_at();

-- 10. Триггеры: обновление search_vector
CREATE FUNCTION fn_events_search_vector() RETURNS trigger LANGUAGE plpgsql AS $$
BEGIN
  NEW.search_vector :=
    setweight(to_tsvector('russian', coalesce((NEW.i18n->'ru'->>'title'),'') ), 'A') ||
    setweight(to_tsvector('russian', coalesce((NEW.i18n->'ru'->>'desc'),'')  ), 'B');
  RETURN NEW;
END; $$;
CREATE TRIGGER trg_events_search_vector
  BEFORE INSERT OR UPDATE ON events
  FOR EACH ROW EXECUTE FUNCTION fn_events_search_vector();

-- 11. Фоновая очистка (pg_cron)
SELECT cron.schedule(
  'daily_cleanup_events',
  '0 0 * * *',
  $$
    WITH moved AS (
      DELETE FROM events
       WHERE ends_at IS NOT NULL AND ends_at < now()
       RETURNING *
    )
    INSERT INTO events_archive
    SELECT * FROM moved;
  $$
);
