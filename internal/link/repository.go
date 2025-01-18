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
func(r *Repository)DeleteLink( link *Link){
	
}
func(r *Repository)GoToLink( link *Link){

}