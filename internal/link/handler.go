package link

import (
	"net/http"

	"github.com/Asker231/link-shortener.git/pkg/req"
)


type LinkHandler struct{
	repo Repository
}

func NewHandlerLink(app *http.ServeMux,repo Repository){
	link := LinkHandler{
			repo: repo,
	}
	app.HandleFunc("POST /link",link.Create())
	app.HandleFunc("PATH /link/{id}",link.Update())
	app.HandleFunc("DELETE /link/{id}",link.Delete())
	app.HandleFunc("GET /link/{hash}",link.GoTo())
}	

func(l *LinkHandler)Create()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		body,err :=req.HandleBody[CreateRequest](w,r)
		if err != nil{
			return
		}
		link := NewLink(body.Url)

		l.repo.CreateLink(link)
	}
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