package db

import (
	"github.com/Asker231/link-shortener.git/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func ConnectDataBase(config config.AppConfig)*gorm.DB{
	
	db , err := gorm.Open(postgres.Open(config.DSN),&gorm.Config{})
	if err != nil{
		panic(err)
	}
	return db

}