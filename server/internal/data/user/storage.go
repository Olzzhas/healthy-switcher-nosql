package user

import (
	"context"
	"server/internal/data/token"
)

type Storage interface {
	Create(ctx context.Context, user *User) (string, error)
	FindOne(ctx context.Context, id string) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error

	CreateOrder(ctx context.Context, user User, order Order) error
	UpdateForToken(ctx context.Context, user User, token token.Token) error
	FindOneByEmail(ctx context.Context, email string) (User, error)
	FindForActivation(ctx context.Context, activationToken string) (User, error)
}
