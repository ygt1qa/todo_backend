package app

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ygt1qa/todo_backend/internal/infrastructures/gqlgen"
	"github.com/ygt1qa/todo_backend/internal/infrastructures/gqlgen/generated"

	"github.com/gin-gonic/gin"
	"github.com/ygt1qa/todo_backend/internal/infrastructures/orm"
	"github.com/ygt1qa/todo_backend/internal/interface/adapter"
)

// Handler todo tasks api
func Handler(r *gin.Engine) {

	resttaskAdapter := adapter.RestNewTaskAdapter(orm.NewOrmHandler())
	v1 := r.Group("v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1.POST("/tasks", func(c *gin.Context) { resttaskAdapter.Create(c) })
		v1.GET("/tasks", func(c *gin.Context) { resttaskAdapter.FetchAll(c) })
		v1.DELETE("/tasks/:id", func(c *gin.Context) { resttaskAdapter.Delete(c) })
		v1.PUT("/tasks/:id", func(c *gin.Context) { resttaskAdapter.Update(c) })

		// Graphql
		v1.POST("/query", graphqlHandler())
		v1.GET("/", playgroundHandler())
	}
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &gqlgen.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
