package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	//module info routes
	router.HandlerFunc(http.MethodPost, "/v1/module-info", app.createModuleInfoHandler)
	router.HandlerFunc(http.MethodGet, "/v1/module-info/{id}", app.getModuleInfoHandler)
	router.HandlerFunc(http.MethodPut, "/v1/module-info/{id}", app.editModuleInfoHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/module-info/{id}", app.deleteModuleInfoHandler)

	//department info routes
	router.HandlerFunc(http.MethodPost, "/v1/department-info", app.createDepInfoHandler)
	router.HandlerFunc(http.MethodGet, "/v1/department-info/{id}", app.getDepInfoHandler)

	//user info routes

	return app.recoverPanic(app.rateLimit(router))
}
