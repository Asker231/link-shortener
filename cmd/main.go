package main

import (
	"net/http"

	"github.com/Asker231/link-shortener.git/config"
	"github.com/Asker231/link-shortener.git/internal/auth"
	"github.com/Asker231/link-shortener.git/internal/link"
	"github.com/Asker231/link-shortener.git/pkg/db"
)

func main() {
	//globalCongif
	appConf, authConf := config.NewConfig()
	
	//connectToDataBase
	_ = db.ConnectDataBase(*appConf)

	//create router app
	app := http.NewServeMux()

	//create serverConfigs
	server := http.Server{
		Addr:    ":8080",
		Handler: app,
	}

	//authHandler
	auth.NewAuthHandler(app,authConf)
	//linkHandler
	link.NewHandlerLink(app)


	server.ListenAndServe()

}
