package main

import "net/http"


func main(){
	app := http.NewServeMux()


	server := http.Server{
		Addr: ":3002",
		Handler: app,
	}

	server.ListenAndServe()
}