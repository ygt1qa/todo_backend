package datastore

import (
	"github.com/ygt1qa/todo_backend/internal/domains/models"
)

// TaskRepository Tasks's Repository Sub
type TaskRepository struct {
	OrmHandler
}

// Store save task
func (t *TaskRepository) Store(m models.Tasks) error {
	err := t.Create(m)
	return err
}
