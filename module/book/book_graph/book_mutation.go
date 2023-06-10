package book_graph

import (
	"context"
	"xyu-graphql/internal/dao/model"
)

type MutationResolver struct{}

// CreateBook is the resolver for the createBook field.
func (r *MutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	if &input == nil {
		return &model.Book{
			ID:   2147483649,
			Name: "book01",
			Author: &model.Author{
				ID:   1002,
				Name: "xyu",
			},
			UUID: "xyu-xxxxwsfnkwv-xwvjihjk-o0n12enjn",
		}, nil
	}
	return &model.Book{
		ID:   2147483649,
		Name: input.Name,
		Author: &model.Author{
			ID:   2147483669,
			Name: "xyu",
		},
		UUID: "uuid-xxxx-xxxxxxxxxxxx",
	}, nil
}

// DeleteBook is the resolver for the createBook field.
func (r *MutationResolver) DeleteBook(ctx context.Context, input model.BookID) (bool, error) {

	return true, nil
}
