// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import "xyu-graphql/internal/dao/model_gen"

type BookMutationObject struct {
	CreateBook *model_gen.Book `json:"createBook"`
	DeleteBook *bool           `json:"deleteBook,omitempty"`
}

type TestQueryObject struct {
	Test *model_gen.Test `json:"test"`
}

type BookQueryObject struct {
	Books       []*model_gen.Book `json:"books"`
	GetBookByID []*model_gen.Book `json:"getBookByID"`
}

type TodoMutationObject struct {
	NewTodo *model_gen.Todo `json:"newTodo"`
}

type TodoQueryObject struct {
	Todos           []*model_gen.Todo `json:"todos"`
	GetTodoByID     []*model_gen.Todo `json:"getTodoByID"`
	GetTodoByUserID []*model_gen.Todo `json:"getTodoByUserID"`
}

type UserMutationObject struct {
	RegisterUser *model_gen.User `json:"registerUser"`
}

type UserQueryObject struct {
	UserList    []*model_gen.User `json:"userList"`
	GetUserByID []*model_gen.User `json:"getUserByID"`
}
