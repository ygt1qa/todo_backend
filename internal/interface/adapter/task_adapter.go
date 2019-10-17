package adapter

import (
	"strconv"

	"github.com/ygt1qa/todo_backend/internal/domains/models"
	"github.com/ygt1qa/todo_backend/internal/interface/datastore"
	"github.com/ygt1qa/todo_backend/internal/usecases"
)

// TaskAdapter struct
type TaskAdapter struct {
	Interactor usecases.TaskInteractor
}

// NewTaskAdapter create new taskadapter
func NewTaskAdapter(orm datastore.OrmHandler) *TaskAdapter {
	return &TaskAdapter{
		Interactor: usecases.TaskInteractor{
			TaskRepository: &datastore.TaskRepository{
				OrmHandler: orm,
			},
		},
	}
}

// Create Add task
func (adapter *TaskAdapter) Create(c Context) {
	t := models.Task{}
	c.Bind(&t)
	err := adapter.Interactor.Add(t)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, "success")
}

// FetchAll get all tasks
func (adapter *TaskAdapter) FetchAll(c Context) {
	result, err := adapter.Interactor.FindAll()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, result)
}

// Delete delete task
func (adapter *TaskAdapter) Delete(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := adapter.Interactor.Remove(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, "success")
}

// Update update task
func (adapter *TaskAdapter) Update(c Context) {
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
