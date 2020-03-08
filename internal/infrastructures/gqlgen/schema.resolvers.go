package gqlgen

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ygt1qa/todo_backend/internal/domains/models"
	"github.com/ygt1qa/todo_backend/internal/infrastructures/gqlgen/generated"
	"github.com/ygt1qa/todo_backend/internal/infrastructures/gqlgen/model"
	"github.com/ygt1qa/todo_backend/internal/infrastructures/orm"
	"github.com/ygt1qa/todo_backend/internal/interface/adapter"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*models.Task, error) {
	taskAdapter := adapter.GqlNewTaskAdapter(orm.NewOrmHandler())
	return taskAdapter.Create(input)
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.EditTodo) (*models.Task, error) {
	taskAdapter := adapter.GqlNewTaskAdapter(orm.NewOrmHandler())
	return taskAdapter.Update(input)
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input int) (*models.Task, error) {
	taskAdapter := adapter.GqlNewTaskAdapter(orm.NewOrmHandler())
	return taskAdapter.Delete(input)
}

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Task, error) {
	taskAdapter := adapter.GqlNewTaskAdapter(orm.NewOrmHandler())
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return taskAdapter.FetchAll(gc)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}
