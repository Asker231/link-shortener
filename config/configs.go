package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


type(
	AppConfig struct{
		DSN string 
	}
	AuthConfig struct{
		TOKEN string `json:"token"`
	}
)


func NewConfig()(appConf *AppConfig,authConf *AuthConfig){
	 err := godotenv.Load("../.env")
	 if err != nil{
		fmt.Println(err.Error())
	 }	
	 return &AppConfig{DSN: os.Getenv("DSN")} , &AuthConfig{TOKEN: os.Getenv("TOKEN")}
}