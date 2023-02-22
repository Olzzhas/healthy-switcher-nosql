package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"server/internal/data/dish"
	dishDb "server/internal/data/dish/db"
	"server/internal/validator"
)

func (app *application) showDishHandler(w http.ResponseWriter, r *http.Request) {
	//var input struct {
	//	id string
	//}

	// НЕ доделано!!!

	id := "63f3d8d33535f36eb46a65ad"

	storage := dishDb.NewStorage(app.mongoClient, "dishes")

	dishRes, err := storage.FindOne(context.Background(), id)

	fmt.Print(reflect.TypeOf(dishRes))

	err = app.writeJSON(w, http.StatusOK, envelope{"dish": dishRes}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createDishHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string `json:"title"`
		Img         string `json:"img"`
		Price       int64  `json:"price"`
		Description string `json:"description"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	var comments []dish.Comment
	var comment dish.Comment
	comments = append(comments, comment)

	item := &dish.Dish{
		Title:       input.Title,
		Img:         input.Img,
		Price:       input.Price,
		Description: input.Description,
		Comments:    comments,
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

func (app *application) getDishesHandler(w http.ResponseWriter, r *http.Request) {
	storage := dishDb.NewStorage(app.mongoClient, "dishes")

	dishes, err := storage.FindAll(r.Context())
	if err != nil {
		log.Fatal(err)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"dishes": dishes}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		DishId      string `json:"dish_id"bson:"dish_id"`
		UserId      string `json:"user_id"bson:"user_id"`
		CommentBody string `json:"comment_body"bson:"comment_body"`
		Rating      int64  `json:"rating"bson:"rating"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	storage := dishDb.NewStorage(app.mongoClient, "dishes")

	dishRes, err := storage.FindOne(context.Background(), input.DishId)
	if err != nil {
		log.Fatal(err)
	}

	var comment dish.Comment

	comment.UserID = input.UserId
	comment.CommentBody = input.CommentBody
	comment.Rating = input.Rating

	err = storage.CreateComment(r.Context(), dishRes, comment)
	if err != nil {
		log.Fatal(err)
	}
}
