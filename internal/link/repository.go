package link

import (
	"gorm.io/gorm"
)


type Repository struct{
	DtaBase *gorm.DB
}

func NewLinkRepository(db *gorm.DB)*Repository{
	return &Repository{
		DtaBase: db,
	}
}

func(r *Repository)GetAll()[]Link{
	 var links []Link
	 r.DtaBase.Find(&links)
	 return links	
}

func(r *Repository)CreateLink( link *Link){
      r.DtaBase.Create(link)
}
func(r *Repository)UpdateLink( link *Link){
	
}
func(r *Repository)DeleteLink(id int)(error){
	 var link Link
	 result := 	r.DtaBase.Delete(&link,"id = ?",id)
	 if result.Error != nil{
		panic(result.Error.Error())
	 }	
	 return nil
}

func(r *Repository)GetLinkById(id int)(*Link,error){
	var link Link
	 result := r.DtaBase.First(&link,"id = ?",id)
	 if result.Error != nil{
		panic(result.Error.Error())
	 }	
	 return &link,nil
}