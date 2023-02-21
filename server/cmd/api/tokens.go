package main

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"log"
	"server/internal/data/token"
	tokenDB "server/internal/data/token/db"
	"time"
)

func (app *application) generateToken(userID string, ttl time.Duration, scope string) (*token.Token, error) {

	token := &token.Token{
		UserID: userID,
		Expiry: time.Now().Add(ttl),
		Scope:  scope,
	}

	randomBytes := make([]byte, 16)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:]
	return token, nil
}

func (app *application) New(userID string, ttl time.Duration, scope string) (*token.Token, error) {
	token, err := app.generateToken(userID, ttl, scope)
	if err != nil {
		return nil, err
	}

	storage := tokenDB.NewStorage(app.mongoClient, "tokens")

	err = storage.Create(context.Background(), *token)
	if err != nil {
		log.Fatal(err)
	}

	return token, err
}
