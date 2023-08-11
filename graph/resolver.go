package graph

import (
	"time"

	"github.com/KazukiHayase/graphql-todo-app-server/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func NewResolver() *Resolver {
	return &Resolver{
		users: []*model.User{
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

type Resolver struct {
	users []*model.User
}

func (r *Resolver) wait() {
	time.Sleep(1 * time.Second)
}
