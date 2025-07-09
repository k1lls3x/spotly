package main

import (
	"net/http"

	"github.com/rs/zerolog/log"

	"spotly/internal/config"
	"spotly/internal/repository"
	"spotly/internal/service"
	httpserver "spotly/internal/server/http"
)

func main() {
	// 1. Настроить логгер
	config.SetupLogger()

	// 2. Подключение к БД (и загрузка .env)
	db := config.Init()

	// 3. Логгер (можно сделать через log.With() если нужно)
	logger := log.Logger

	// 4. Репозитории
	caffeRepo := repository.NewCaffeRepository(db, logger)
	fsRepo := repository.NewFederalSubjectRepository(db, logger) // если используешь

	// 5. Сервисы
	caffeSvc := service.NewCaffeService(caffeRepo, logger)
	fsSvc := service.NewFederalSubjectService(fsRepo, logger)

	// 6. Роутер
	router := httpserver.NewRouter(logger, fsSvc, caffeSvc)

	// 7. Запуск сервера
	addr := ":8080"
	logger.Info().Str("addr", addr).Msg("🚀 HTTP-сервер запущен")
	if err := http.ListenAndServe(addr, router); err != nil {
		logger.Fatal().Err(err).Msg("❌ Сервер завершил работу с ошибкой")
	}
}
