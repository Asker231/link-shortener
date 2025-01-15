package req

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func IsValid[T any](payload T)error{
	validate := validator.New()
	err := validate.Struct(payload) 
	if err != nil{
		fmt.Println(err.Error())
	}
	return err
}