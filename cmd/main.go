package main

import (
	"log/slog"
	"net/http"
	"os"
)


type app struct{
	logger *slog.Logger

	
}

func main(){
	logger:= slog.New(slog.NewTextHandler(os.Stdout,nil))

	app:= &app{
		logger: logger,
	}

	


	logger.Info("server is running on port :5132")

	http.ListenAndServe(":5132",app.routes())


}