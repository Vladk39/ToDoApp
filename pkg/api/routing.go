package api

import (
	_ "todoapi/pkg/api/docs"
	"todoapi/pkg/config"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"go.uber.org/zap"
)

func StartServer(h *Handler, c *config.Config, logger *zap.Logger) {
	app := fiber.New()

	tasks := app.Group("/tasks")

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	tasks.Post("/", h.CreateTask)
	tasks.Get("/", h.GetAllTasks)
	tasks.Put("/:id", h.UpdateTask)
	tasks.Delete("/:id", h.DeleteTask)

	if err := app.Listen(c.ServerPort); err != nil {
		logger.Error("Ошибка при запуске сервера", zap.Error(err))
	}
}
