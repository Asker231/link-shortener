package req

import (
	"net/http"
	"github.com/Asker231/link-shortener.git/pkg/res"
)
	func HandleBody[T any](w http.ResponseWriter, r *http.Request)(*T,error){
		body,err := DecodeJson[T](r.Body)
		if err != nil{
			res.Response(w,err,402)
			return nil,err
		}
		err = IsValid(body)
		if err != nil{
			res.Response(w,err.Error(),402)
			return nil,err
		}
		return body,nil
	}