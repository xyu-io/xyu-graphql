package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"xyu-graphql/internal/dao"
	"xyu-graphql/internal/dao/model"
	"xyu-graphql/internal/graph_server/generated"
	book "xyu-graphql/module/book/api_graph"
	todo "xyu-graphql/module/todo/api_graph"
)

// TodoServer is the resolver for the TodoServer field.
func (r *mutationResolver) TodoServer(ctx context.Context) (*dao.TodoMutationObject, error) {
	return &dao.TodoMutationObject{}, nil
}

// BookServer is the resolver for the BookServer field.
func (r *mutationResolver) BookServer(ctx context.Context) (*dao.BookMutationObject, error) {
	return &dao.BookMutationObject{}, nil
}

// TodoServer is the resolver for the TodoServer field.
func (r *queryResolver) TodoServer(ctx context.Context) (*dao.TodoQueryObject, error) {
	return &dao.TodoQueryObject{}, nil
}

// BookServer is the resolver for the BookServer field.
func (r *queryResolver) BookServer(ctx context.Context) (*dao.BookQueryObject, error) {
	return &dao.BookQueryObject{}, nil
}

// CreateBook is the resolver for the createBook field.
func (r *bookMutationObjectResolver) CreateBook(ctx context.Context, obj *dao.BookMutationObject, input model.NewBook) (*model.Book, error) {
	resolver := book.MutationResolver{}
	return resolver.CreateBook(ctx, input)
}

// Books is the resolver for the books field.
func (r *bookQueryObjectResolver) Books(ctx context.Context, obj *dao.BookQueryObject) ([]*model.Book, error) {
	resolver := book.QueryResolver{}
	return resolver.Books(ctx)
}

// GetBookByID is the resolver for the getBookByID field.
func (r *bookQueryObjectResolver) GetBookByID(ctx context.Context, obj *dao.BookQueryObject, input model.BookID) ([]*model.Book, error) {
	resolver := book.QueryResolver{}
	return resolver.GetBookByID(ctx, input)
}

// NewTodo is the resolver for the newTodo field.
func (r *todoMutationObjectResolver) NewTodo(ctx context.Context, obj *dao.TodoMutationObject, input model.NewTodo) (*model.Todo, error) {
	resolver := todo.MutationResolver{}
	return resolver.NewTodo(ctx, input)
}

// Todos is the resolver for the todos field.
func (r *todoQueryObjectResolver) Todos(ctx context.Context, obj *dao.TodoQueryObject) ([]*model.Todo, error) {
	resolver := todo.QueryResolver{}
	return resolver.Todos(ctx)
}

// BookMutationObject returns generated.BookMutationObjectResolver implementation.
func (r *Resolver) BookMutationObject() generated.BookMutationObjectResolver {
	return &bookMutationObjectResolver{r}
}

// BookQueryObject returns generated.BookQueryObjectResolver implementation.
func (r *Resolver) BookQueryObject() generated.BookQueryObjectResolver {
	return &bookQueryObjectResolver{r}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// TodoMutationObject returns generated.TodoMutationObjectResolver implementation.
func (r *Resolver) TodoMutationObject() generated.TodoMutationObjectResolver {
	return &todoMutationObjectResolver{r}
}

// TodoQueryObject returns generated.TodoQueryObjectResolver implementation.
func (r *Resolver) TodoQueryObject() generated.TodoQueryObjectResolver {
	return &todoQueryObjectResolver{r}
}

type bookMutationObjectResolver struct{ *Resolver }
type bookQueryObjectResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoMutationObjectResolver struct{ *Resolver }
type todoQueryObjectResolver struct{ *Resolver }
