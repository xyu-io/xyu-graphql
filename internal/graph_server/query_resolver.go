package graph_server

import (
	"context"
	"xyu-graphql/internal/dao"
	main "xyu-graphql/internal/dao/model"
	book "xyu-graphql/module/book/api_graph"
	todo "xyu-graphql/module/todo/api_graph"
)

type QueryObjectResolver struct{}

// Books is the resolver for the books field.
func (r *QueryObjectResolver) Books(ctx context.Context, obj *dao.BookQueryObject) ([]*main.Book, error) {
	resolver := book.QueryResolver{}
	return resolver.Books(ctx)
}

// Todos is the resolver for the todos field.
func (r *QueryObjectResolver) Todos(ctx context.Context, obj *dao.TodoQueryObject) ([]*main.Todo, error) {
	resolver := todo.QueryResolver{}
	return resolver.Todos(ctx)
}
