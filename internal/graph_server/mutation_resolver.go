package graph_server

import (
	"context"
	"xyu-graphql/internal/dao"
	main "xyu-graphql/internal/dao/model"
	book "xyu-graphql/module/book/api_graph"
	todo "xyu-graphql/module/todo/api_graph"
)

type MutationObjectResolver struct{}

// CreateBook is the resolver for the createBook field.
func (r *MutationObjectResolver) CreateBook(ctx context.Context, _ *dao.BookMutationObject, input main.NewBook) (*main.Book, error) {
	resolver := book.MutationResolver{}
	return resolver.CreateBook(ctx, input)
}

// DeleteBook is the resolver for the createBook field.
func (r *MutationObjectResolver) DeleteBook(ctx context.Context, _ *dao.BookMutationObject, input main.BookID) (*bool, error) {
	resolver := book.MutationResolver{}
	return resolver.DeleteBook(ctx, input)
}

// NewTodo is the resolver for the newTodo field.
func (r *MutationObjectResolver) NewTodo(ctx context.Context, _ *dao.TodoMutationObject, input main.NewTodo) (*main.Todo, error) {
	resolver := todo.MutationResolver{}
	return resolver.NewTodo(ctx, input)
}
