package usecases

import (
	"github.com/ygt1qa/todo_backend/internal/domains/models"
)

// TaskInteractor task
type TaskInteractor struct {
	TaskRepository TaskRepository
}

// Add save task
func (i *TaskInteractor) Add(t models.Task) error {
	return i.TaskRepository.Store([]models.Task{t})
}
