package graph_server

import (
	"context"
	"xyu-graphql/internal/dao"
	"xyu-graphql/internal/graph_server/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

/**
 * -----------------------------------Mutation--------------------------------------
 */

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver {
	return &QueryResolver{}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver {
	return &MutationResolver{}
}

/**
 * --------------------------------config server-----------------------------------------
 */

type (
	QueryResolver    struct{}
	MutationResolver struct{}
)

func (r *QueryResolver) TodoServer(ctx context.Context) (*dao.TodoQueryObject, error) {
	return &dao.TodoQueryObject{}, nil
}

func (r *QueryResolver) BookServer(ctx context.Context) (*dao.BookQueryObject, error) {
	return &dao.BookQueryObject{}, nil
}

func (r *MutationResolver) TodoServer(ctx context.Context) (*dao.TodoMutationObject, error) {
	return &dao.TodoMutationObject{}, nil
}

func (r *MutationResolver) BookServer(ctx context.Context) (*dao.BookMutationObject, error) {
	return &dao.BookMutationObject{}, nil
}

/**
 * --------------------------------book server-----------------------------------------
 */

// BookMutationObject returns generated.BookMutationObjectResolver implementation.
func (r *Resolver) BookMutationObject() generated.BookMutationObjectResolver {
	return &MutationObjectResolver{}
}

// BookQueryObject returns generated.BookQueryObjectResolver implementation.
func (r *Resolver) BookQueryObject() generated.BookQueryObjectResolver {
	return &QueryObjectResolver{}
}

/**
 * --------------------------------todos server-----------------------------------------
 */

// TodoMutationObject returns generated.TodoMutationObjectResolver implementation.
func (r *Resolver) TodoMutationObject() generated.TodoMutationObjectResolver {
	return &MutationObjectResolver{}
}

// TodoQueryObject returns generated.TodoQueryObjectResolver implementation.
func (r *Resolver) TodoQueryObject() generated.TodoQueryObjectResolver {
	return &QueryObjectResolver{}
}
