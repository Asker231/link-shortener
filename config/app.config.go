package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)



type (
	AppConfig struct{
		DB *DataBase
		AuthCnf *AuthConf
	}
	DataBase struct{
		DNS string
	}

	AuthConf struct{
		Secret string
	}
)


func NewAppConfig()*AppConfig{
	if err := godotenv.Load(".env"); err != nil{
		log.Panic(err.Error())
		return nil
	}	
	return &AppConfig{
		DB: &DataBase{
			DNS: os.Getenv("DNS"),
		},
		AuthCnf: &AuthConf{
			Secret: os.Getenv("SECRET"),
		},
	}
}