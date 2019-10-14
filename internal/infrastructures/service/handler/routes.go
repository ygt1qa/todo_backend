package handler

import (
	"fmt"

	"github.com/ygt1qa/todo_backend/internal/infrastructures/service/app"

	"github.com/gin-gonic/gin"
)

// NewRouter func
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(cors())

	app.Handler(r)

	return r
}

// via https://github.com/gin-contrib/cors/issues/29#issuecomment-397859488
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, DELETE, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS sent")
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
