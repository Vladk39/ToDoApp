package api

import (
	"fmt"
	"strconv"
	"todoapi/pkg/taskRepository"
	"todoapi/pkg/taskService"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	s *taskService.TaskService
}

func NewHandler(s *taskService.TaskService) *Handler {
	return &Handler{
		s: s,
	}
}

func ParseID(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

// TaskCreateRequest - структура запроса для создания или обновления задачи
// @Description Структура запроса для создания или обновления задачи
type TaskCreateRequest struct {
	// Заголовок задачи
	Title string `json:"title"`
	// Описание задачи
	Description string `json:"description"`
	// Статус задачи
	Status string `json:"status"`
}

// CreateTask добавляет новую задачу.
// @Summary Создание задачи
// @Description Добавить новую задачу
// @Tags tasks
// @Accept json
// @Produce plain
// @Param task body TaskCreateRequest true "Новая задача"
// @Success 200 {string} string "Задача создана"
// @Failure 400 {string} string "неверный формат запроса"
// @Failure 500 {string} string "ошибка добавления задачи"
// @Router /tasks [post]
func (h *Handler) CreateTask(c *fiber.Ctx) error {
	var t taskRepository.Task
	if err := c.BodyParser(&t); err != nil {
		return c.Status(400).SendString("неверный формат запроса")
	}

	err := h.s.CreateTaskService(c.Context(), t)
	if err != nil {
		return c.Status(500).SendString("ошибка добавления задачи")
	}

	return c.Status(200).SendString(fmt.Sprintf("Задача создана: %s", t.Title))
}

// DeleteTask удаляет задачу по ID.
// @Summary Удаление задачи
// @Description Удалить задачу по ID
// @Tags tasks
// @Produce plain
// @Param id path int true "ID задачи"
// @Success 200 {string} string "Задача удалена"
// @Failure 400 {string} string "неверный ID"
// @Failure 500 {string} string "Ошибка удаления"
// @Router /tasks/{id} [delete]
func (h *Handler) DeleteTask(c *fiber.Ctx) error {
	idstr := c.Params("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.Status(400).SendString("неверный ID")
	}

	if err := h.s.DeleteTaskService(c.Context(), id); err != nil {
		return c.Status(500).SendString("Ошибка удаления")
	}

	return c.Status(200).SendString(fmt.Sprintf("Задача удалена, ID: %d", id))
}

// UpdateTask обновляет задачу.
// @Summary Обновление задачи
// @Description Обновить задачу по ID
// @Tags tasks
// @Accept json
// @Produce plain
// @Param id path int true "ID задачи"
// @Param task body TaskCreateRequest true "Обновлённая задача"
// @Success 200 {string} string "Задача обновлена"
// @Failure 400 {string} string "Неверный формат"
// @Failure 500 {string} string "Ошибка обновления задачи"
// @Router /tasks/{id} [put]
func (h *Handler) UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var t taskRepository.Task
	if err := c.BodyParser(&t); err != nil {
		return c.Status(400).SendString("Неверный формат")
	}

	t.ID = ParseID(id)
	err := h.s.UpdateTaskService(c.Context(), t)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления задачи")
	}

	return c.Status(200).SendString(fmt.Sprintf("Задача %s обновлена", t.Title))
}

// GetAllTasks возвращает все задачи.
// @Summary Получить все задачи
// @Description Получить список всех задач
// @Tags tasks
// @Produce json
// @Success 200 {array} taskRepository.Task
// @Failure 500 {string} string "Ошибка получения списка задач"
// @Router /tasks [get]
func (h *Handler) GetAllTasks(c *fiber.Ctx) error {
	tasks, err := h.s.GetAllTasksService(c.Context())
	if err != nil {
		return c.Status(500).SendString("Ошибка получения списка задач")
	}

	return c.JSON(tasks)
}
