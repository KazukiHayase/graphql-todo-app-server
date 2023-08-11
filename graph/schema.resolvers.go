package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/KazukiHayase/graphql-todo-app-server/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	r.wait()
	var user *model.User
	for _, u := range r.users {
		if u.ID == input.UserID {
			tmp := u
			user = &tmp
		}
	}

	if user == nil {
		return nil, fmt.Errorf("User not found")
	}

	todo := &model.Todo{
		ID:   fmt.Sprintf("T%d", len(user.Todos)+1),
		Text: input.Text,
		Done: false,
	}

	user.Todos = append(user.Todos, todo)

	return todo, nil
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, id string, input model.UpdateTodo) (*model.Todo, error) {
	r.wait()
	var todo *model.Todo
	for _, u := range r.users {
		for _, t := range u.Todos {
			if t.ID == id {
				tmp := *t
				todo = &tmp
			}
		}
	}

	if todo == nil {
		return nil, fmt.Errorf("Todo not found")
	}

	if input.Text != nil {
		todo.Text = *input.Text
	}
	if input.Done != nil {
		todo.Done = *input.Done
	}

	return todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	r.wait()
	var todos []*model.Todo
	for _, u := range r.users {
		for _, t := range u.Todos {
			tmp := *t
			todos = append(todos, &tmp)
		}
	}

	return todos, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	r.wait()
	var user *model.User
	for _, u := range r.users {
		if u.ID == id {
			tmp := u
			user = &tmp
		}
	}

	return user, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
