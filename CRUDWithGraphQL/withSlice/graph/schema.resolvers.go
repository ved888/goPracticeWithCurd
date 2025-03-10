package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"withSlice/graph/model"

	"github.com/vektah/gqlparser/gqlerror"
)

var Users []model.User

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, id string, name string, address string) (*model.User, error) {
	var user model.User
	// Validate input payload
	if strings.TrimSpace(id) == "" || strings.TrimSpace(name) == "" || strings.TrimSpace(address) == "" {
		return nil, &gqlerror.Error{
			Message: "Invalid input: ID, Name, and Address are required",
			Extensions: map[string]interface{}{
				"code": http.StatusBadRequest,
			},
		}
	}
	user.ID = id
	user.Name = name
	user.Address = address
	Users = append(Users, user)
	return &user, nil
}

// UpdateUser is the resolver for the UpdateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, name string, address string) (*model.User, error) {
	// Validate input payload
	if strings.TrimSpace(name) == "" || strings.TrimSpace(address) == "" {
		return nil, &gqlerror.Error{
			Message: "Invalid input: ID, Name, and Address are required",
			Extensions: map[string]interface{}{
				"code": http.StatusBadRequest, // 400 Bad Request
			},
		}
	}
	// Find the user in global storage
	for i, u := range Users {
		if u.ID == id {
			// Update user data
			Users[i].Name = name
			Users[i].Address = address

			return &Users[i], nil
		}
	}
	return nil, &gqlerror.Error{
		Message: fmt.Sprintf("User with ID %s not found", id),
		Extensions: map[string]interface{}{
			"code": http.StatusNotFound, // 404 Not Found
		},
	}
}

// RemoveUser is the resolver for the removeUser field.
func (r *mutationResolver) RemoveUser(ctx context.Context, input model.DeleteUser) (*model.User, error) {
	var deletedUser *model.User
	newUsers := []model.User{} // Slice of non-pointer `User`

	for _, user := range Users {
		if user.ID == input.ID {
			// Take the address of `user` to assign to `deletedUser`
			deletedUser = &user
			continue // Skip adding this user to newUsers
		}
		newUsers = append(newUsers, user)
	}

	// If no user was found, return a 404 error
	if deletedUser == nil {
		return nil, &gqlerror.Error{
			Message: fmt.Sprintf("User with ID %s not found", input.ID),
			Extensions: map[string]interface{}{
				"code": http.StatusNotFound, // 404 Not Found
			},
		}
	}

	// Update the global Users slice
	Users = newUsers

	return deletedUser, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	// Convert []model.User to []*model.User
	users := make([]*model.User, len(Users))
	for i := range Users {
		users[i] = &Users[i]
	}

	return users, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	for i := range Users {
		if Users[i].ID == id {
			return &Users[i], nil // Return a pointer to the actual object in Users
		}
	}
	// Return a GraphQL error with HTTP Status Code 404
	return nil, &gqlerror.Error{
		Message: fmt.Sprintf("User with ID %s not found", id),
		Extensions: map[string]interface{}{
			"code": http.StatusNotFound, // 404 Not Found
		},
	}
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	var Users []model.User
*/
