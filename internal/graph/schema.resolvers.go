package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/abeja-project/abeja/internal/graph/generated"
	"github.com/abeja-project/abeja/internal/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo, err := r.saveTodo(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("creating todo: %w", err)
	}

	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := r.getTodos(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting todos: %w", err)
	}
	return todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
