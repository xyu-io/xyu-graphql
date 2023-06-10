package book_graph

import (
	"context"
	"xyu-graphql/internal/dao/model"
)

type QueryResolver struct{}

// Books is the resolver for the books field.
func (r *QueryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	bk := &model.Book{
		ID:   9223372036854775807,
		Name: "xyu-graphql-example",
		Author: &model.Author{
			ID:   9023372036854775807,
			Name: "xyu",
		},
		UUID: "xyu-xcwv-graphql-vnvw10001-1002ssvsd",
	}
	var bks []*model.Book
	bks = append(bks, bk)
	return bks, nil
}

func (r *QueryResolver) GetBookByID(ctx context.Context, id model.BookID) ([]*model.Book, error) {
	bk := &model.Book{
		ID:   id.ID,
		Name: "xyu-graphql-example",
		Author: &model.Author{
			ID:   9023372036854775807,
			Name: "xyu",
		},
		UUID: "xyu-xcwv-graphql-vnvw10001-1002ssvsd",
	}
	var bks []*model.Book
	bks = append(bks, bk)
	return bks, nil
}
