package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	//router errors
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	//test
	router.HandlerFunc(http.MethodGet, "/api/healthcheck", app.requireAuthenticatedUser(app.healthcheckHandler))

	//dish
	router.HandlerFunc(http.MethodGet, "/api/showDish", app.showDishHandler)
	router.HandlerFunc(http.MethodPost, "/api/dish", app.createDishHandler)
	router.HandlerFunc(http.MethodGet, "/api/topDishes", app.getDishesHandler)

	//comment
	router.HandlerFunc(http.MethodPut, "/api/comment", app.requireActivatedUser(app.createCommentHandler))

	//user
	router.HandlerFunc(http.MethodPost, "/api/user", app.registerUserHandler)
	router.HandlerFunc(http.MethodPost, "/api/user/byToken", app.findByTokenHandler)
	router.HandlerFunc(http.MethodPut, "/api/users/activated", app.activateUserHandler)
	router.HandlerFunc(http.MethodPost, "/api/tokens/authentication", app.createAuthenticationTokenHandler)

	//order
	router.HandlerFunc(http.MethodPost, "/api/order", app.requireActivatedUser(app.createOrderHandler))

	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}
