package orm

import (
	"context"
	"database/sql"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/ygt1qa/todo_backend/internal/domains/models"
	"github.com/ygt1qa/todo_backend/internal/interface/datastore"

	// use mysql
	_ "github.com/go-sql-driver/mysql"
)

var ormhandler OrmHandler

// OrmHandler is settings
type OrmHandler struct {
	db boil.ContextExecutor
}

// InitDB sql connection init
func InitDB() boil.ContextExecutor {

	// connect to db
	db, err := sql.Open("mysql", "mysql:mysql@tcp(db:3306)/mysql_test")
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	// validate whether or not the connection string was correct
	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}

// NewOrmHandler create ormhandler
func NewOrmHandler() datastore.OrmHandler {
	return new(OrmHandler)
}

// Create stores a one or more todos into DB
func (o *OrmHandler) Create(m models.Tasks) error {
	o.db = InitDB()

	// insert for each todo
	for _, v := range m {
		var task *Task
		task = TaskToBoilTask(v)
		if err := task.Insert(context.Background(), o.db, boil.Infer()); err != nil {
			return err
		}
	}
	return nil
}

// Remove delete task
func (o *OrmHandler) Remove(id int) error {
	o.db = InitDB()

	task, _ := FindTask(context.Background(), o.db, int64(id))
	_, err := task.Delete(context.Background(), o.db)
	return err
}

// UpdateByID update task
func (o *OrmHandler) UpdateByID(id int, t models.Task) error {
	o.db = InitDB()

	task, _ := FindTask(context.Background(), o.db, int64(id))
	task.Name = t.Name
	task.Description = null.StringFrom(t.Description)
	_, err := task.Update(context.Background(), o.db, boil.Infer())
	if err != nil {
		return err
	}
	return err
}

// FindAll get all tasks
func (o *OrmHandler) FindAll() ([]*models.Task, error) {
	o.db = InitDB()
	result, err := Tasks().All(context.Background(), o.db)

	TaskList := []*models.Task{}
	for _, value := range result {
		TaskList = append(TaskList, BoilTaskToTask(value))
	}
	return TaskList, err
}

// TaskToBoilTask entity task conversion sql boiler entity
func TaskToBoilTask(t models.Task) *Task {
	return &Task{
		Name:        t.Name,
		Description: null.StringFrom(t.Description),
	}
}

// BoilTaskToTask sql boiler entity conversion entity task
func BoilTaskToTask(t *Task) *models.Task {
	return &models.Task{
		ID:          int(t.ID),
		Name:        t.Name,
		Description: t.Description.String,
	}
}
