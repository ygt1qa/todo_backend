package datastore

import (
	"github.com/ygt1qa/todo_backend/internal/domains/models"
)

type OrmHandler interface {
	Create(models.Tasks) error
	FindAll() ([]*models.Task, error)
	Remove(int) error
}
