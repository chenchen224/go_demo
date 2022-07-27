package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"context"

	"gitlab.com/chenxk/demo/graph/model"
)

type Resolver struct{}

func (r *Resolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return &model.Todo{}, nil
}

// Todos is the resolver for the todos field.
func (r *Resolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return []*model.Todo{}, nil
}
