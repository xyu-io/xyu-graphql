package todo_graph

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

func (r *QueryResolver) GetTodoByID(ctx context.Context, id *model.TodoID) ([]*model.Todo, error) {
	bk := &model.Todo{
		ID:   *id.ID,
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

func (r *QueryResolver) GetTodoByUserID(ctx context.Context, userID *model.UserID) ([]*model.Todo, error) {
	bk := &model.Todo{
		ID:   100010,
		Text: "读书",
		Done: true,
		User: &model.User{
			ID:   userID.UserID,
			Name: "张三",
		},
	}
	var bks []*model.Todo
	bks = append(bks, bk)
	return bks, nil
}
