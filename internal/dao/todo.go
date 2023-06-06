package dao

import "xyu-graphql/internal/dao/model"

type TodoMutationObject struct {
	NewTodo *model.Todo `json:"newTodo"`
}

type TodoQueryObject struct {
	Todos []*model.Todo `json:"todos"`
}
