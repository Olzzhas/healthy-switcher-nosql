package db

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"server/internal/data/dish"
)

type db struct {
	collection *mongo.Collection
	logger     *log.Logger
}

func (d *db) Create(ctx context.Context, dish dish.Dish) (string, error) {
	result, err := d.collection.InsertOne(ctx, dish)
	if err != nil {
		return "", fmt.Errorf("failed to create user due to error: %v", err)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}

	return "", fmt.Errorf("failed to convert objectid to hex. probably oid: %s", oid)
}

func (d *db) FindOne(ctx context.Context, id string) (dish dish.Dish, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dish, fmt.Errorf("failed to convert hex to objectid. hex: %s", id)
	}

	filter := bson.M{"_id": oid}

	result := d.collection.FindOne(ctx, filter)

	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			//TODO ErrEntityNotFound
			return dish, fmt.Errorf("ErrEntityNotFound")
		}
		return dish, fmt.Errorf("failed to find one user by id: %s due to error: %v", id, err)
	}

	if err = result.Decode(&dish); err != nil {
		return dish, fmt.Errorf("failed to decode user(id: %s) from db due to error: %v", id, err)
	}

	return dish, nil
}

func (d *db) Update(ctx context.Context, dish dish.Dish) error {
	objectID, err := primitive.ObjectIDFromHex(dish.ID)
	if err != nil {
		return fmt.Errorf("failed to convert user ID to ObjectID. Id=%s", dish.ID)
	}

	filter := bson.M{"_id": objectID}

	userBytes, err := bson.Marshal(dish)
	if err != nil {
		return fmt.Errorf("failed to marshal user. error: %v", err)
	}

	var updateUserObj bson.M
	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal userbytes: %v", err)
	}

	delete(updateUserObj, "_id")

	update := bson.M{
		"$set": updateUserObj,
	}

	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to execute update user query. error: %v", err)
	}

	if result.MatchedCount == 0 {
		//TODO ErrEntityNotFound
		return fmt.Errorf("not found")
	}

	fmt.Printf("Matched %d documents and Modified %d documents", result.MatchedCount, result.ModifiedCount)

	return nil
}

func (d *db) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert user ID to ObjectID. ID=%s", id)
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

func NewStorage(database *mongo.Database, collection string) dish.Storage {

	return &db{
		collection: database.Collection(collection),
	}
}
