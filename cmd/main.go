package main

import (
	"log"
	"net/http"

	"github.com/Asker231/link-shortener.git/config"
	"github.com/Asker231/link-shortener.git/internal/auth"
)



func main() {
	_, authConf :=config.NewConfig()
	app := http.NewServeMux()
	server := http.Server{
		Addr: ":30",
		Handler: app,
	}

	auth.NewAuthHandker(app,authConf)


	log.Fatal(server.ListenAndServe())
 
}
