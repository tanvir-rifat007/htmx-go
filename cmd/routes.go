package main

import "net/http"


func (app *app) routes()http.Handler{

	mux:= http.NewServeMux()

	mux.HandleFunc("GET /",app.home)
	mux.HandleFunc("POST /contacts",app.contacts)

	return app.loggerMiddleware(mux)

}