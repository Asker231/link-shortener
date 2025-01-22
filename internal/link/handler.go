package link

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Asker231/link-shortener.git/pkg/req"
	"github.com/Asker231/link-shortener.git/pkg/res"
)


type LinkHandler struct{
	repo Repository
}

func NewHandlerLink(app *http.ServeMux,repo Repository){
	link := LinkHandler{
			repo: repo,
	}
	app.HandleFunc("GET /link/all",link.All())
	app.HandleFunc("POST /link",link.Create())
	app.HandleFunc("PATH /link/{id}",link.Update())
	app.HandleFunc("DELETE /link/{id}",link.Delete())
	app.HandleFunc("GET /link/{id}",link.Getlink())
}	

func(l *LinkHandler)All()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		allLinks := l.repo.GetAll()
		fmt.Println(allLinks)
		res.Response(w,allLinks,200)
	}

}
func(l *LinkHandler)Create()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		body,err :=req.HandleBody[CreateRequest](w,r)
		if err != nil{
			return
		}
		link := NewLink(body.Url)
		l.repo.CreateLink(link)

		res.Response(w,"success",200)
	}
}
func(l *LinkHandler)Delete()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		pv := r.PathValue("id")
		id ,err := strconv.Atoi(pv)
		if err != nil{
			fmt.Println(err.Error())
		} 
		err = l.repo.DeleteLink(id)
		if err != nil{
			fmt.Println(err.Error())
		}
		links := l.repo.GetAll()
		res.Response(w,links,200)

	}
}
func(l *LinkHandler)Update()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r.PathValue("id")
	}
}
func(l *LinkHandler)Getlink()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		pv := r.PathValue("id")
		id ,err := strconv.Atoi(pv)
		if err != nil{
			fmt.Println(err.Error())
		}
		link ,err := l.repo.GetLinkById(id)
		if err != nil{
			fmt.Println(err)
		}
		res.Response(w,link,200)

	}
}