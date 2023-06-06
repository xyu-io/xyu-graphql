package api_graph

import (
	"context"
	"xyu-graphql/internal/dao/model"
)

type MutationResolver struct{}

// NewTodo is the resolver for the createBook field.
func (r *MutationResolver) NewTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	if &input == nil {
		return &model.Todo{
			ID:   0,
			Text: "xxx",
			Done: false,
			User: nil,
		}, nil
	}
	return &model.Todo{
		ID:   1,
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID:   2,
			Name: "xxs",
		},
	}, nil
}
