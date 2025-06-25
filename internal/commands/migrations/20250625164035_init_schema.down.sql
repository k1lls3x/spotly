-- 1. Удаляем расписание pg_cron
SELECT cron.unschedule('daily_cleanup_events');

-- 2. Удаляем триггеры обновления updated_at
DROP TRIGGER IF EXISTS trg_cities_updated_at ON cities;
DROP TRIGGER IF EXISTS trg_categories_updated_at ON categories;
DROP TRIGGER IF EXISTS trg_events_updated_at ON events;

-- 3. Удаляем триггер обновления search_vector
DROP TRIGGER IF EXISTS trg_events_search_vector ON events;

-- 4. Удаляем функции
DROP FUNCTION IF EXISTS fn_set_updated_at();
DROP FUNCTION IF EXISTS fn_events_search_vector();

-- 5. Удаляем таблицы (в обратном порядке создания)
DROP TABLE IF EXISTS events_archive;
DROP TABLE IF EXISTS entity_media;
DROP TABLE IF EXISTS media;
DROP TABLE IF EXISTS event_categories;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS cities;
DROP TABLE IF EXISTS timezones;
DROP TABLE IF EXISTS federal_subjects;

-- 6. (Опционально) Удаляем расширения
DROP EXTENSION IF EXISTS pg_cron;
DROP EXTENSION IF EXISTS pg_trgm;
DROP EXTENSION IF EXISTS pgcrypto;
