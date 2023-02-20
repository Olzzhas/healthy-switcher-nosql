package main

import (
	"fmt"
	"net/http"
	"server/internal/data/dish"
	"time"
)

func (app *application) showDishHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	fmt.Print(id)

	dish := dish.Dish{
		CreatedAt:   time.Now(),
		Title:       "Plov",
		Description: "Vkusno",
		Rating:      4,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"dish": dish}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
