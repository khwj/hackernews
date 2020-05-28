package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/khwj/hackernews/graph/generated"
	"github.com/khwj/hackernews/graph/model"
	"github.com/khwj/hackernews/internal/links"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	var link links.Link
	// var user model.User
	link.URL = input.URL
	link.Description = input.Description
	// user.Name = "test"
	// link.User = &user
	id := link.Save()
	return &model.Link{ID: strconv.FormatInt(id, 10), Description: link.Description, URL: link.URL}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	dbLinks := links.GetAll()
	var links []*model.Link
	for _, link := range dbLinks {
		newLink := model.Link{ID: link.ID, Description: link.Description, URL: link.URL}
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
