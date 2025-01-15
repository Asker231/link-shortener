package auth

import (
	"net/http"
	"github.com/Asker231/link-shortener.git/config"
)

type(
	Auth struct{
		config *config.AuthConfig
	}
)

func NewAuthHandker(app *http.ServeMux,authConf *config.AuthConfig){
	auth := &Auth{
		config: authConf,
	}	
	app.HandleFunc("POST /auth/register",auth.Register())
	app.HandleFunc("POST /auth/login",auth.Login())
}

func(a *Auth)Register()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
	     //прочитать body и провалидировать 
		//реализовать функцию возврата JSON	
	}
}
func(a *Auth)Login()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

