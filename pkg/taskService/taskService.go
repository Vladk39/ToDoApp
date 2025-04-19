package taskService

import (
	"context"
	"todoapi/pkg/config"
	"todoapi/pkg/taskRepository"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type TaskService struct {
	repo   *taskRepository.Repository
	c      *config.Config
	logger *zap.Logger
}

func NewTaskService(repo *taskRepository.Repository, c *config.Config, logger *zap.Logger) *TaskService {
	return &TaskService{
		repo:   repo,
		c:      c,
		logger: logger,
	}
}

func (s *TaskService) CreateTaskService(ctx context.Context, t taskRepository.Task) error {
	err := s.repo.AddTask(ctx, t)
	if err != nil {
		s.logger.Error("Ошибка добавления задачи", zap.Error(err))
		return errors.Wrap(err, "ошибка добавления задачи")
	}
	return nil
}

func (s *TaskService) DeleteTaskService(ctx context.Context, id int) error {
	err := s.repo.DeleteTask(ctx, id)
	if err != nil {
		s.logger.Error("Ошибка  удаления задачи", zap.Error(err))
		return errors.Wrap(err, "ошибка добавления задачи")
	}

	return nil
}

func (s *TaskService) UpdateTaskService(ctx context.Context, t taskRepository.Task) error {
	err := s.repo.UpdateTask(ctx, t)
	if err != nil {
		s.logger.Error("Ошибка обновления задачи", zap.Error(err))
		return errors.Wrap(err, "ошибка обновления задачи")
	}

	return nil
}

func (s *TaskService) GetAllTasksService(ctx context.Context) ([]taskRepository.Task, error) {
	tasks, err := s.repo.GetAllTasks(ctx)
	if err != nil {
		s.logger.Error("Ошибка получения списка задач", zap.Error(err))
		return nil, errors.Wrap(err, "ошибка получения списка задач")
	}

	return tasks, nil
}
