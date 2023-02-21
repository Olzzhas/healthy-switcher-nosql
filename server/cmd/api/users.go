package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	dishDB "server/internal/data/dish/db"
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

func (app *application) createOrderHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		dish_id string
		user_id string
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	dishStorage := dishDB.NewStorage(app.mongoClient, "dishes")

	dish, err := dishStorage.FindOne(r.Context(), input.dish_id)
	if err != nil {
		log.Fatal(err)
	}

	storage := userDB.NewStorage(app.mongoClient, "users")

	candidate, err := storage.FindOne(r.Context(), input.user_id)
	if err != nil {
		log.Fatal(err)
	}

	var order user.Order
	order.Dish = dish

	storage.CreateOrder(r.Context(), candidate, order)

}
