package main

import (
	"net/http"

	"github.com/Asker231/link-shortener.git/internal/auth"
)


func main(){

	app := http.NewServeMux()
	
	auth.NewAuth(app)

	server := http.Server{
		Addr: ":3002",
		Handler: app,
	}

	server.ListenAndServe()
}