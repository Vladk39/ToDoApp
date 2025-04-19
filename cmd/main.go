package main

import (
	"todoapi/pkg/api"
	"todoapi/pkg/config"
	"todoapi/pkg/taskRepository"
	"todoapi/pkg/taskService"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// @title ToDo API
// @version 1.0
// @description Простой API для управления задачами

// @host localhost:8080
// @BasePath /
func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		errors.Wrap(err, "ошибка запуска логера")
		return
	}
	defer logger.Sync()
	logger.Info("Приложение запущено")

	cfg, err := config.GetConfig()
	if err != nil {
		logger.Error("не удалось получить конфигурации", zap.Error(err))
		return
	}

	repo, err := taskRepository.NewRepository(cfg)
	if err != nil {
		logger.Error("Не удалось создать репозиторий", zap.Error(err))
		return
	}
	logger.Info("Репозиторий создан")

	service := taskService.NewTaskService(repo, cfg, logger)

	handler := api.NewHandler(service)

	api.StartServer(handler, cfg, logger)
}
