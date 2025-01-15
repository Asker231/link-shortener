package main

import (
	"net/http"

	"github.com/Asker231/link-shortener.git/config"
	"github.com/Asker231/link-shortener.git/internal/auth"
)

func main() {
	_, authConf := config.NewConfig()
	app := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: app,
	}

	auth.NewAuthHandker(app, authConf)

	server.ListenAndServe()

}
