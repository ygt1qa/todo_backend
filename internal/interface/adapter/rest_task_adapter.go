package adapter

import (
	"strconv"

	"github.com/ygt1qa/todo_backend/internal/domains/models"
	"github.com/ygt1qa/todo_backend/internal/interface/datastore"
	"github.com/ygt1qa/todo_backend/internal/usecases"
)

// RestTaskAdapter struct
type RestTaskAdapter struct {
	Interactor usecases.TaskInteractor
}

// RestNewTaskAdapter create new taskadapter
func RestNewTaskAdapter(orm datastore.OrmHandler) *RestTaskAdapter {
	return &RestTaskAdapter{
		Interactor: usecases.TaskInteractor{
			TaskRepository: &datastore.TaskRepository{
				OrmHandler: orm,
			},
		},
	}
}

// Create Add task
func (adapter *RestTaskAdapter) Create(c Context) {
	t := models.Task{}
	c.Bind(&t)
	task, err := adapter.Interactor.Add(t)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, task)
}

// FetchAll get all tasks
func (adapter *RestTaskAdapter) FetchAll(c Context) {
	result, err := adapter.Interactor.FindAll()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, result)
}

// Delete delete task
func (adapter *RestTaskAdapter) Delete(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := adapter.Interactor.Remove(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, task)
}

// Update update task
func (adapter *RestTaskAdapter) Update(c Context) {
	t := models.Task{}
	c.Bind(&t)
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := adapter.Interactor.UpdateById(id, t)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, task)
}
