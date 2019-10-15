package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ygt1qa/todo_backend/internal/infrastructures/orm"
	"github.com/ygt1qa/todo_backend/internal/interface/adapter"
)

// Handler todo tasks api
func Handler(r *gin.Engine) {

	taskAdapter := adapter.NewTaskAdapter(orm.NewOrmHandler())
	v1 := r.Group("v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1.POST("/tasks", func(c *gin.Context) { taskAdapter.Create(c) })
		v1.GET("/tasks", func(c *gin.Context) { taskAdapter.FetchAll(c) })
		v1.GET("/tasks/:id")
		v1.DELETE("/tasks/:id", func(c *gin.Context) { taskAdapter.Delete(c) })
		v1.PUT("/tasks/:id")
	}
}
