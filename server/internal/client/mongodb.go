package client

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, database string) (db *mongo.Database, err error) {
	mongoDBURL := "mongodb+srv://olzhas:olzhas04@cluster0.imwcna0.mongodb.net/?retryWrites=true&w=majority"

	// Connect

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDBURL))
	if err != nil {
		return nil, fmt.Errorf("failed to create new client due to error: %v", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb due to error: %v", err)
	}

	//Ping

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping to mongodb due to error: %v", err)
	}

	return client.Database(database), nil
}
