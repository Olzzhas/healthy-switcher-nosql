package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"server/internal/data/token"
)

type db struct {
	collection *mongo.Collection
	logger     *log.Logger
}

func (d *db) Create(ctx context.Context, token token.Token) error {
	result, err := d.collection.InsertOne(ctx, token)
	if err != nil {
		return fmt.Errorf("failed to create token due to error: %v", err)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return nil
	}

	return fmt.Errorf("failed to convert objectid to hex. probably oid: %s", oid)
}

func (d *db) FindOne(ctx context.Context, id string) (token.Token, error) {
	//TODO implement me
	panic("implement me")
}

func (d *db) Update(ctx context.Context, token token.Token) error {
	//TODO implement me
	panic("implement me")
}

// потом исправлю
func (d *db) Delete(ctx context.Context, id string) error {
	//query := `
	//	DELETE FROM tokens
	//	WHERE scope = $1 AND user_id = $2
	//`
	//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//defer cancel()
	//_, err := m.DB.ExecContext(ctx, query, scope, userID)
	//return err

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert token ID to ObjectID. ID=%s", id)
	}

	filter := bson.M{"_id": objectID}

	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("not found")
	}

	fmt.Printf("Deleted %d documents", result.DeletedCount)

	return nil
}

func NewStorage(database *mongo.Database, collection string) token.Storage {

	return &db{
		collection: database.Collection(collection),
	}
}
