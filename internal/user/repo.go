package user

import "gorm.io/gorm"


type UserRepo struct{
	DataBase *gorm.DB
}


//create user

//find user