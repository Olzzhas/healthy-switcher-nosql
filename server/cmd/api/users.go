package main

import (
	"context"
	"fmt"
	"net/http"
	"server/internal/data/user"
	userDB "server/internal/data/user/db"
	"server/internal/validator"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"bson:"name"`
		Email    string `json:"email"bson:"email"`
		Password string `json:"password"bson:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	candidate := &user.User{
		Name:      input.Name,
		Email:     input.Email,
		Activated: false,
	}

	err = candidate.Password.Set(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	v := validator.New()

	if user.ValidateUser(v, candidate); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	storage := userDB.NewStorage(app.mongoClient, "users")

	userRes, err := storage.Create(context.Background(), candidate)

	fmt.Fprintf(w, userRes)

}
