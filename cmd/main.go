package main

import (
	"log/slog"
	"net/http"
	"os"
	"sync"
)


type app struct{
	logger *slog.Logger
  Contacts []contact
	mu sync.Mutex

	
}

type contact struct{
	Name string
	Email string
}

func main(){
	logger:= slog.New(slog.NewTextHandler(os.Stdout,nil))

	app:= &app{
		logger: logger,
	}

	


	logger.Info("server is running on port :5132")

	http.ListenAndServe(":5132",app.routes())


}