package user

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"server/internal/data/token"
	"server/internal/data/user"
)

type db struct {
	collection *mongo.Collection
	logger     *log.Logger
}

func (d *db) Create(ctx context.Context, user *user.User) (string, error) {
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create user due to error: %v", err)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}

	return "", fmt.Errorf("failed to convert objectid to hex. probably oid: %s", oid)
}

func (d *db) FindOne(ctx context.Context, id string) (u user.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert hex to objectid. hex: %s", id)
	}

	filter := bson.M{"_id": oid}

	result := d.collection.FindOne(ctx, filter)

	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			//TODO ErrEntityNotFound
			return u, fmt.Errorf("ErrEntityNotFound")
		}
		return u, fmt.Errorf("failed to find one user by id: %s due to error: %v", id, err)
	}

	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user(id: %s) from db due to error: %v", id, err)
	}

	return u, nil
}

func (d *db) Update(ctx context.Context, user user.User) error {
	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return fmt.Errorf("failed to convert user ID to ObjectID. Id=%s", user.ID)
	}

	filter := bson.M{"_id": objectID}

	userBytes, err := bson.Marshal(user)
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

func (d *db) CreateOrder(ctx context.Context, user user.User, order user.Order) error {
	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return fmt.Errorf("failed to convert user ID to ObjectID. Id=%s", user.ID)
	}

	filter := bson.M{"_id": objectID}

	userBytes, err := bson.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user. error: %v", err)
	}

	var updateUserObj bson.M
	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal userbytes: %v", err)
	}

	delete(updateUserObj, "_id")

	update := bson.D{{"$push", bson.D{{"orders", order}}}}

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

func (d *db) UpdateForToken(ctx context.Context, user user.User, token token.Token, scope string) error {
	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return fmt.Errorf("(update)failed to convert user ID to ObjectID. Id=%s", user.ID)
	}

	filter := bson.M{"_id": objectID}

	userBytes, err := bson.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user. error: %v", err)
	}

	var updateUserObj bson.M
	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal userbytes: %v", err)
	}

	delete(updateUserObj, "_id")

	update := bson.D{{"$set", bson.D{{scope, token}}}}

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

func (d *db) FindOneByEmail(ctx context.Context, email string) (u user.User, err error) {
	filter := bson.M{"email": email}

	result := d.collection.FindOne(ctx, filter)

	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			//TODO ErrEntityNotFound
			return u, fmt.Errorf("ErrEntityNotFound")
		}
		return u, fmt.Errorf("failed to find one user by email: %s due to error: %v", email, err)
	}

	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user(email: %s) from db due to error: %v", email, err)
	}

	return u, nil
}

func (d *db) FindForActivation(ctx context.Context, activationToken string) (u user.User, err error) {
	filter := bson.M{"activation.token": activationToken}

	result := d.collection.FindOne(ctx, filter)

	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			//TODO ErrEntityNotFound
			return u, fmt.Errorf("ErrEntityNotFound")
		}
		return u, fmt.Errorf("failed to find one user by activation token: %s due to error: %v", activationToken, err)
	}

	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user(activation token: %s) from db due to error: %v", activationToken, err)
	}

	return u, nil
}

func (d *db) FindForAuthentication(ctx context.Context, authenticationToken string) (u user.User, err error) {
	filter := bson.M{"authentication.token": authenticationToken}

	result := d.collection.FindOne(ctx, filter)

	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			//TODO ErrEntityNotFound
			return u, fmt.Errorf("ErrEntityNotFound")
		}
		return u, fmt.Errorf("failed to find one user by authentication token: %s due to error: %v", authenticationToken, err)
	}

	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user(authentication token: %s) from db due to error: %v", authenticationToken, err)
	}

	return u, nil
}

func NewStorage(database *mongo.Database, collection string) user.Storage {

	return &db{
		collection: database.Collection(collection),
	}
}
