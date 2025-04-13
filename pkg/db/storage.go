package db

import (
	"log"

	"github.com/Asker231/link-shortener.git/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Storage struct{
	DataBase *gorm.DB
}

func NewStorage(cnf *config.DataBase)*Storage{
	conn,err := gorm.Open(postgres.Open(cnf.DNS),&gorm.Config{})
	if err != nil{
		log.Fatal(err.Error())
		return nil
	}
	return &Storage{
		DataBase: conn,
	}
}