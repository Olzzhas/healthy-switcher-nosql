package main

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"log"
	"net/http"
	"server/internal/data/token"
	tokenDB "server/internal/data/token/db"
	"server/internal/data/user"
	userDB "server/internal/data/user/db"
	"server/internal/validator"
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

	storage := tokenDB.NewStorage(app.mongoClient, scope)

	err = storage.Create(context.Background(), *token)
	if err != nil {
		log.Fatal(err)
	}

	return token, err
}

func (app *application) createAuthenticationTokenHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	user.ValidateEmail(v, input.Email)
	user.ValidatePasswordPlaintext(v, input.Password)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	storage := userDB.NewStorage(app.mongoClient, "users")

	user, err := storage.FindOneByEmail(context.Background(), input.Email)
	if err != nil {
		log.Fatal(err)
	}

	match, err := user.Password.Matches(input.Password, user.Hash)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if !match {
		app.invalidCredentialsResponse(w, r)
		return
	}

	token, err := app.New(user.ID, 24*time.Hour, token.ScopeAuthentication)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	candidateForToken, err := storage.FindOneByEmail(context.Background(), user.Email)
	if err != nil {
		log.Fatal(err)
	}

	err = storage.UpdateForToken(context.Background(), candidateForToken, *token, "authentication")
	if err != nil {
		log.Fatal(err)
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"authentication_token": token}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
