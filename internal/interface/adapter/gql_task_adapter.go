package adapter

import (
	"github.com/ygt1qa/todo_backend/internal/domains/models"
	"github.com/ygt1qa/todo_backend/internal/infrastructures/gqlgen/model"
	"github.com/ygt1qa/todo_backend/internal/interface/datastore"
	"github.com/ygt1qa/todo_backend/internal/usecases"
)

// TaskAdapter struct
type GqlTaskAdapter struct {
	Interactor usecases.TaskInteractor
}

// GqlNewTaskAdapter create new taskadapter
func GqlNewTaskAdapter(orm datastore.OrmHandler) *GqlTaskAdapter {
	return &GqlTaskAdapter{
		Interactor: usecases.TaskInteractor{
			TaskRepository: &datastore.TaskRepository{
				OrmHandler: orm,
			},
		},
	}
}

// Create Add task
func (adapter *GqlTaskAdapter) Create(input model.NewTodo) (*models.Task, error) {
	t := models.Task{
		Name:        input.Name,
		Description: input.Description,
	}

	return adapter.Interactor.Add(t)
}

// FetchAll get all tasks
func (adapter *GqlTaskAdapter) FetchAll(c Context) ([]*models.Task, error) {
	return adapter.Interactor.FindAll()
}

// Delete delete task
func (adapter *GqlTaskAdapter) Delete(id int) (*models.Task, error) {
	return adapter.Interactor.Remove(id)
}

// Update update task
func (adapter *GqlTaskAdapter) Update(input model.EditTodo) (*models.Task, error) {
	t := models.Task{
		Name:        input.Name,
		Description: input.Description,
	}
	return adapter.Interactor.UpdateById(input.ID, t)
}
