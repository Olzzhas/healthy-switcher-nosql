package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"server/internal/data/token"
)

type db struct {
	collection *mongo.Collection
	logger     *log.Logger
}

func (d *db) Create(ctx context.Context, token token.Token) (*token.Token, error) {
	panic("implement me")
}

func (d *db) FindOne(ctx context.Context, id string) (token.Token, error) {
	//TODO implement me
	panic("implement me")
}

func (d *db) Update(ctx context.Context, token token.Token) error {
	//TODO implement me
	panic("implement me")
}

func (d *db) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewStorage(database *mongo.Database, collection string) token.Storage {

	return &db{
		collection: database.Collection(collection),
	}
}
