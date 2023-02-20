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
	router.HandlerFunc(http.MethodGet, "/api/dish/:id", app.showDishHandler)

	//user

	return router
}
