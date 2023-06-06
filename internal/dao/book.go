package dao

import "xyu-graphql/internal/dao/model"

type BookMutationObject struct {
	CreateBook *model.Book `json:"createBook"`
	DeleteBook *bool       `json:"deleteBook,omitempty"`
}

type BookQueryObject struct {
	Books []*model.Book `json:"books"`
}
