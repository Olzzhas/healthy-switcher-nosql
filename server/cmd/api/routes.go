package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	//router errors
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	//test
	router.HandlerFunc(http.MethodGet, "/api/healthcheck", app.healthcheckHandler)

	//dish
	router.HandlerFunc(http.MethodGet, "/api/showDish", app.showDishHandler)
	router.HandlerFunc(http.MethodPost, "/api/dish", app.createDishHandler)
	router.HandlerFunc(http.MethodGet, "/api/topDishes", app.getDishesHandler)

	//comment
	router.HandlerFunc(http.MethodPut, "/api/comment", app.createCommentHandler)

	//user
	router.HandlerFunc(http.MethodPost, "/api/user", app.registerUserHandler)

	//order
	router.HandlerFunc(http.MethodPost, "/api/order", app.createOrderHandler)

	return router
}
