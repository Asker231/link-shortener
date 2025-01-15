package auth

import (
	"fmt"
	"net/http"

	"github.com/Asker231/link-shortener.git/config"
	"github.com/Asker231/link-shortener.git/pkg/req"
	"github.com/Asker231/link-shortener.git/pkg/res"
)

type (
	Auth struct {
		config *config.AuthConfig
	}
)

func NewAuthHandker(app *http.ServeMux, authConf *config.AuthConfig) {
	auth := &Auth{
		config: authConf,
	}
	app.HandleFunc("POST /auth/register", auth.Register())
	app.HandleFunc("POST /auth/login", auth.Login())
}

func (a *Auth) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := req.HandleBody[RegisterRequest](w, r)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(payload)
		response := RegisterResponse{
			Token: a.config.TOKEN,
		}
		res.Response(w, response, 200)

	}
}
func (a *Auth) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
