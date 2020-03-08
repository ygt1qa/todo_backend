package datastore

import (
	"github.com/ygt1qa/todo_backend/internal/domains/models"
)

type OrmHandler interface {
	Create(models.Tasks) (*models.Task, error)
	FindAll() ([]*models.Task, error)
	Remove(int) (*models.Task, error)
	UpdateByID(int, models.Task) (*models.Task, error)
}
