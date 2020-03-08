package gqlgen

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

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
