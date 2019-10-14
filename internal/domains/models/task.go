package models

// Task TODO task
type Task struct {
	ID          int
	Name        string
	Description string
}

// Tasks Task objs
type Tasks []Task
