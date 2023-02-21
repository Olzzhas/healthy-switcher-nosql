package token

import (
	"context"
)

type Storage interface {
	Create(ctx context.Context, token Token) error
	FindOne(ctx context.Context, id string) (Token, error)
	Update(ctx context.Context, token Token) error
	Delete(ctx context.Context, id string) error
}
