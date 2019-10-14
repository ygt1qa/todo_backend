package main

import (
	"github.com/ygt1qa/todo_backend/internal/infrastructures/service/handler"
)

func main() {
	router := handler.NewRouter()
	router.Run(":8000")
}
