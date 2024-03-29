package usecases

import "github.com/ygt1qa/todo_backend/internal/domains/models"

// TaskRepository for interface
type TaskRepository interface {
	Store(models.Tasks) (*models.Task, error)
	GetAll() ([]*models.Task, error)
	Erase(int) (*models.Task, error)
	Update(int, models.Task) (*models.Task, error)
}
