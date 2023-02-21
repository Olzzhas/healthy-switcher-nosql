package dish

import (
	"context"
)

type Storage interface {
	Create(ctx context.Context, dish *Dish) (string, error)
	FindOne(ctx context.Context, id string) (Dish, error)
	Update(ctx context.Context, dish Dish) error
	Delete(ctx context.Context, id string) error

	FindAll(ctx context.Context) (dish []Dish, err error)
	CreateComment(ctx context.Context, dish Dish, comment Comment) error
}
