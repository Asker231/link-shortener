package main

import (
	"net/http"

	"github.com/Asker231/link-shortener.git/config"
	"github.com/Asker231/link-shortener.git/internal/auth"
	"github.com/Asker231/link-shortener.git/internal/link"
	"github.com/Asker231/link-shortener.git/pkg/db"
)

func main() {
	appConf, authConf := config.NewConfig()
	_ = db.ConnectDataBase(*appConf)
	app := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: app,
	}


	auth.NewAuthHandler(app,authConf)
	//linkHandler
	link.NewHandlerLink(app)
	server.ListenAndServe()

}
