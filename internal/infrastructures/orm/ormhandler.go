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
	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/mysql_test")
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

// TaskToBoilTask entity task conversion sql boiler entity
func TaskToBoilTask(t models.Task) *Task {
	return &Task{
		Name:        t.Name,
		Description: null.StringFrom(t.Description),
	}
}
