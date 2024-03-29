package usecases

import (
	"github.com/ygt1qa/todo_backend/internal/domains/models"
)

// TaskInteractor task
type TaskInteractor struct {
	TaskRepository TaskRepository
}

// Add save task
func (i *TaskInteractor) Add(t models.Task) (*models.Task, error) {
	return i.TaskRepository.Store([]models.Task{t})
}

// FindAll get all tasks
func (i *TaskInteractor) FindAll() ([]*models.Task, error) {
	return i.TaskRepository.GetAll()
}

func (i *TaskInteractor) Remove(id int) (*models.Task, error) {
	return i.TaskRepository.Erase(id)
}

func (i *TaskInteractor) UpdateById(id int, t models.Task) (*models.Task, error) {
	return i.TaskRepository.Update(id, t)
}
