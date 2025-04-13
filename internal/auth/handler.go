package auth

import "net/http"

type AuthHandler struct{

}

func NewAuth(app *http.ServeMux){
	a := &AuthHandler{}
	app.HandleFunc("GET /",a.Register())
}

func(a *AuthHandler)Register()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from go"))
	}
}

func(a *AuthHandler)Login()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {}
}