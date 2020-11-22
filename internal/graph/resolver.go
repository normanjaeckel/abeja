package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/abeja-project/abeja/internal/graph/model"
)

// Database does ... (TODO)
type Database interface {
	GetObjects(ctx context.Context, model string) (map[int]json.RawMessage, error)
	NewID(ctx context.Context, model string) (int, error)
	UpdateObject(ctx context.Context, model string, id int, data json.RawMessage) error
}

// Resolver does ... (TODO)
type Resolver struct {
	db Database
}

// NewResolver does ... (TODO)
func NewResolver(db Database) *Resolver {
	return &Resolver{db: db}
}

func (r *Resolver) getUser(ctx context.Context, id string) (*model.User, error) {
	if id != "1" {
		return nil, &userDoesNotExistError{id: id}
	}
	anja := &model.User{
		ID:   "1",
		Name: "Anja",
	}
	return anja, nil
}

func (r *Resolver) getTodos(ctx context.Context) ([]*model.Todo, error) {
	rawTodos, err := r.db.GetObjects(ctx, "todo")
	if err != nil {
		return nil, fmt.Errorf("getting todos: %w", err)
	}

	todos := make([]*model.Todo, 0, len(rawTodos))

	for id, rawTodo := range rawTodos {
		// TODO: test as pointer
		var todo model.Todo
		if err := json.Unmarshal(rawTodo, &todo); err != nil {
			return nil, fmt.Errorf("decoding todo %d: %w", id, err)
		}

		todos = append(todos, &todo)
	}

	return todos, nil
}

func (r *Resolver) saveTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	user, err := r.getUser(ctx, input.UserID)
	if err != nil {
		return nil, fmt.Errorf("getting user: %w", err)
	}

	id, err := r.db.NewID(ctx, "todo")
	if err != nil {
		return nil, fmt.Errorf("getting new todo id: %w", err)
	}

	todo := &model.Todo{
		Text: input.Text,
		Done: false,
		User: user,
		ID:   strconv.Itoa(id),
	}

	rawTodo, err := json.Marshal(todo)
	if err != nil {
		return nil, fmt.Errorf("encoding todo: %w", err)
	}

	if err := r.db.UpdateObject(ctx, "todo", id, rawTodo); err != nil {
		return nil, fmt.Errorf("creating todo: %w", err)
	}

	return todo, nil
}

type userDoesNotExistError struct {
	id string
}

func (e *userDoesNotExistError) Error() string {
	return fmt.Sprintf("User with id %s does not exist", e.id)
}
