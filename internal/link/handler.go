package link

import (
	"net/http"
)


type LinkHandler struct{}

func NewHandlerLink(app *http.ServeMux){
	link := LinkHandler{}
	app.HandleFunc("POST /link",link.Create())
	app.HandleFunc("PATH /link/{id}",link.Update())
	app.HandleFunc("DELETE /link/{id}",link.Delete())
	app.HandleFunc("GET /link/{hash}",link.GoTo())
}	

func(l *LinkHandler)Create()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {}
}
func(l *LinkHandler)Delete()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r.PathValue("id")

	}
}
func(l *LinkHandler)Update()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r.PathValue("id")
	}
}
func(l *LinkHandler)GoTo()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

	}
}