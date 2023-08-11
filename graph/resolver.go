package graph

import "github.com/KazukiHayase/graphql-todo-app-server/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	users []model.User
}

func NewResolver() *Resolver {
	return &Resolver{
		users: []model.User{
			{
				ID:    "1",
				Name:  "User1",
				Todos: []*model.Todo{},
			},
			{
				ID:    "2",
				Name:  "User2",
				Todos: []*model.Todo{},
			},
		},
	}
}
