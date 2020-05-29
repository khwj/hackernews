package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/khwj/hackernews/graph/generated"
	"github.com/khwj/hackernews/graph/model"
	"github.com/khwj/hackernews/internal/auth"
	"github.com/khwj/hackernews/internal/links"
	"github.com/khwj/hackernews/internal/users"
	"github.com/khwj/hackernews/pkg/jwt"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.LinkInput) (*model.Link, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Link{}, fmt.Errorf("Access denied")
	}
	var link links.Link
	link.URL = input.URL
	link.Description = input.Description
	link.User = user
	id := link.Save()
	return &model.Link{
		ID:          strconv.FormatInt(id, 10),
		Description: link.Description,
		URL:         link.URL,
		User:        &model.User{ID: user.ID, Name: user.Username},
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (string, error) {
	user := users.User{Username: input.Username, Password: input.Password}
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	user.Create()
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	user := users.User{Username: input.Username, Password: input.Password}
	if !user.Authenticate() {
		return "", &users.WrongUsernameOrPasswordError{}
	}
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("Access denied")
	}

	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	dbLinks := links.GetAll()
	var links []*model.Link
	for _, link := range dbLinks {
		user := model.User{ID: link.ID, Name: link.User.Username}
		newLink := model.Link{ID: link.ID, Description: link.Description, URL: link.URL, User: &user}
		links = append(links, &newLink)
	}
	return links, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
