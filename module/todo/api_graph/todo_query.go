package api_graph

import (
	"context"
	"xyu-graphql/internal/dao/model"
)

type QueryResolver struct{}

// Todos is the resolver for the books field.
func (r *QueryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	bk := &model.Todo{
		ID:   100010,
		Text: "读书",
		Done: true,
		User: &model.User{
			ID:   10101,
			Name: "张三",
		},
	}
	var bks []*model.Todo
	bks = append(bks, bk)
	return bks, nil
}
