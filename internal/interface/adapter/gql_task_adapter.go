package adapter

import (
	"strconv"

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
func (adapter *GqlTaskAdapter) Delete(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := adapter.Interactor.Remove(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, "success")
}

// Update update task
func (adapter *GqlTaskAdapter) Update(c Context) {
	t := models.Task{}
	c.Bind(&t)
	id, _ := strconv.Atoi(c.Param("id"))
	err := adapter.Interactor.UpdateById(id, t)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, "success")
}
