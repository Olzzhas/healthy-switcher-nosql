package main

import (
	"context"
	"fmt"
	"net/http"
	"server/internal/data/dish"
	dishDb "server/internal/data/dish/db"
	"server/internal/validator"
)

func (app *application) showDishHandler(w http.ResponseWriter, r *http.Request) {
	input_id := "63f3beb8e9053d73256db2c0"

	storage := dishDb.NewStorage(app.mongoClient, "dishes")

	dishRes, err := storage.FindOne(context.Background(), input_id)

	err = app.writeJSON(w, http.StatusOK, envelope{"dish": dishRes}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createDishHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	item := &dish.Dish{
		Title:       input.Title,
		Description: input.Description,
	}

	v := validator.New()

	if dish.ValidateDish(v, item); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	storage := dishDb.NewStorage(app.mongoClient, "dishes")

	dishRes, err := storage.Create(context.Background(), item)

	fmt.Fprintf(w, dishRes)

}
