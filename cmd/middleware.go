package main

import "net/http"


func (app *app) loggerMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter,r *http.Request){
   var (
            ip     = r.RemoteAddr
            proto  = r.Proto
            method = r.Method
            uri    = r.URL.RequestURI()
        )

        app.logger.Info("received request", "ip", ip, "proto", proto, "method", method, "uri", uri)

        next.ServeHTTP(w, r)		 
	})
}