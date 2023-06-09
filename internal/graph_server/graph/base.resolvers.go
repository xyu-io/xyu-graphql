package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"xyu-graphql/internal/dao/model"
	"xyu-graphql/internal/graph_server/generated"
	"xyu-graphql/module/book/book_graph"
	"xyu-graphql/module/todo/todo_graph"
	"xyu-graphql/module/user/user_graph"
)

// CreateBook is the resolver for the createBook field.
func (r *bookMutationObjectResolver) CreateBook(ctx context.Context, obj *model.BookMutationObject, input model.NewBook) (*model.Book, error) {
	resolver := book_graph.MutationResolver{}
	return resolver.CreateBook(ctx, input)
}

// Books is the resolver for the books field.
func (r *bookQueryObjectResolver) Books(ctx context.Context, obj *model.BookQueryObject) ([]*model.Book, error) {
	resolver := book_graph.QueryResolver{}
	return resolver.Books(ctx)
}

// GetBookByID is the resolver for the getBookByID field.
func (r *bookQueryObjectResolver) GetBookByID(ctx context.Context, obj *model.BookQueryObject, input model.BookID) ([]*model.Book, error) {
	resolver := book_graph.QueryResolver{}
	return resolver.GetBookByID(ctx, input)
}

// TodoServer is the resolver for the TodoServer field.
func (r *mutationResolver) TodoServer(ctx context.Context) (*model.TodoMutationObject, error) {
	return &model.TodoMutationObject{}, nil
}

// BookServer is the resolver for the BookServer field.
func (r *mutationResolver) BookServer(ctx context.Context) (*model.BookMutationObject, error) {
	return &model.BookMutationObject{}, nil
}

// UserServer is the resolver for the UserServer field.
func (r *mutationResolver) UserServer(ctx context.Context) (*model.UserMutationObject, error) {
	return &model.UserMutationObject{}, nil
}

// TodoServer is the resolver for the TodoServer field.
func (r *queryResolver) TodoServer(ctx context.Context) (*model.TodoQueryObject, error) {
	return &model.TodoQueryObject{}, nil
}

// BookServer is the resolver for the BookServer field.
func (r *queryResolver) BookServer(ctx context.Context) (*model.BookQueryObject, error) {
	return &model.BookQueryObject{}, nil
}

// UserServer is the resolver for the UserServer field.
func (r *queryResolver) UserServer(ctx context.Context) (*model.UserQueryObject, error) {
	return &model.UserQueryObject{}, nil
}

// Test is the resolver for the Test field.
func (r *queryResolver) Test(ctx context.Context) (*model.TestQueryObject, error) {
	return &model.TestQueryObject{}, nil
}

// Test is the resolver for the test field.
func (r *testQueryObjectResolver) Test(ctx context.Context, obj *model.TestQueryObject) (*model.Test, error) {
	return &model.Test{Title: "test"}, nil
}

// NewTodo is the resolver for the newTodo field.
func (r *todoMutationObjectResolver) NewTodo(ctx context.Context, obj *model.TodoMutationObject, input model.NewTodo) (*model.Todo, error) {
	resolver := todo_graph.MutationResolver{}
	return resolver.NewTodo(ctx, input)
}

// Todos is the resolver for the todos field.
func (r *todoQueryObjectResolver) Todos(ctx context.Context, obj *model.TodoQueryObject) ([]*model.Todo, error) {
	resolver := todo_graph.QueryResolver{}
	return resolver.Todos(ctx)
}

// GetTodoByID is the resolver for the getTodoByID field.
func (r *todoQueryObjectResolver) GetTodoByID(ctx context.Context, obj *model.TodoQueryObject, input *model.TodoID) ([]*model.Todo, error) {
	resolver := todo_graph.QueryResolver{}
	return resolver.GetTodoByID(ctx, input)
}

// GetTodoByUserID is the resolver for the getTodoByUserID field.
func (r *todoQueryObjectResolver) GetTodoByUserID(ctx context.Context, obj *model.TodoQueryObject, input *model.UserID) ([]*model.Todo, error) {
	resolver := todo_graph.QueryResolver{}
	return resolver.GetTodoByUserID(ctx, input)
}

// RegisterUser is the resolver for the registerUser field.
func (r *userMutationObjectResolver) RegisterUser(ctx context.Context, obj *model.UserMutationObject, input model.NewUser) (*model.User, error) {
	resolver := user_graph.MutationResolver{}
	return resolver.RegisterUser(ctx, input)
}

// UserList is the resolver for the userList field.
func (r *userQueryObjectResolver) UserList(ctx context.Context, obj *model.UserQueryObject) ([]*model.User, error) {
	resolver := user_graph.QueryResolver{}
	return resolver.UserList(ctx)
}

// GetUserByID is the resolver for the getUserByID field.
func (r *userQueryObjectResolver) GetUserByID(ctx context.Context, obj *model.UserQueryObject, input model.UserID) ([]*model.User, error) {
	resolver := user_graph.QueryResolver{}
	return resolver.GetUserByID(ctx, input)
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

// TestQueryObject returns generated.TestQueryObjectResolver implementation.
func (r *Resolver) TestQueryObject() generated.TestQueryObjectResolver {
	return &testQueryObjectResolver{r}
}

// TodoMutationObject returns generated.TodoMutationObjectResolver implementation.
func (r *Resolver) TodoMutationObject() generated.TodoMutationObjectResolver {
	return &todoMutationObjectResolver{r}
}

// TodoQueryObject returns generated.TodoQueryObjectResolver implementation.
func (r *Resolver) TodoQueryObject() generated.TodoQueryObjectResolver {
	return &todoQueryObjectResolver{r}
}

// UserMutationObject returns generated.UserMutationObjectResolver implementation.
func (r *Resolver) UserMutationObject() generated.UserMutationObjectResolver {
	return &userMutationObjectResolver{r}
}

// UserQueryObject returns generated.UserQueryObjectResolver implementation.
func (r *Resolver) UserQueryObject() generated.UserQueryObjectResolver {
	return &userQueryObjectResolver{r}
}

type bookMutationObjectResolver struct{ *Resolver }
type bookQueryObjectResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type testQueryObjectResolver struct{ *Resolver }
type todoMutationObjectResolver struct{ *Resolver }
type todoQueryObjectResolver struct{ *Resolver }
type userMutationObjectResolver struct{ *Resolver }
type userQueryObjectResolver struct{ *Resolver }
