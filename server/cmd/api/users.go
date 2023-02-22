package main

import (
	"context"
	"log"
	"net/http"
	dishDB "server/internal/data/dish/db"
	"server/internal/data/token"
	"server/internal/data/user"
	userDB "server/internal/data/user/db"
	"server/internal/validator"
	"time"
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
		Hash:      nil,
		Activated: false,
	}

	passHash, err := candidate.Password.Set(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	candidate = &user.User{
		Name:      input.Name,
		Email:     input.Email,
		Hash:      passHash,
		Activated: false,
	}

	v := validator.New()

	if user.ValidateUser(v, candidate); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	storage := userDB.NewStorage(app.mongoClient, "users")

	userRes, err := storage.Create(context.Background(), candidate)
	if err != nil {
		log.Fatal(err)
	}

	app.writeJSON(w, http.StatusOK, envelope{"userRes": userRes}, nil)
	if err != nil {
		log.Fatal(err)
	}

	token, err := app.New(candidate.ID, 3*24*time.Hour, token.ScopeActivation)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	candidateForToken, err := storage.FindOneByEmail(context.Background(), candidate.Email)
	if err != nil {
		log.Fatal(err)
	}

	err = storage.UpdateForToken(context.Background(), candidateForToken, *token, "activation")
	if err != nil {
		log.Fatal(err)
	}
	//app.background(func() {
	//
	//	data := map[string]any{
	//		"activationToken": token.Plaintext,
	//		"userID":          candidate.ID,
	//	}
	//	err = app.mailer.Send(candidate.Email, "user_welcome.tmpl", data)
	//	if err != nil {
	//		app.logger.PrintError(err, nil)
	//	}
	//})
	//err = app.writeJSON(w, http.StatusAccepted, envelope{"user": candidate}, nil)
	//if err != nil {
	//	app.serverErrorResponse(w, r, err)
	//}

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

func (app *application) activateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		TokenPlaintext string `json:"token"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	if token.ValidateTokenPlaintext(v, input.TokenPlaintext); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	storage := userDB.NewStorage(app.mongoClient, "users")

	user, err := storage.FindForActivation(context.Background(), input.TokenPlaintext)
	if err != nil {
		log.Fatal(err)
	}

	user.Activated = true

	err = storage.Update(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}

	//TODO delete activation token from user document

	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
