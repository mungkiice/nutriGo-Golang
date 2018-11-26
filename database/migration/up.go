package migration

import (
	. "github.com/mungkiice/goNutri/database"
	"github.com/mungkiice/goNutri/model"
	"log"
)

func Up(){
	db := DB.AutoMigrate(
		&model.User{},
		&model.Makanan{},
		&model.Pola{},
		&model.Article{},
		&model.History{})
	if err := db.Error; err != nil {
		log.Println("An Error has occured", err)
	}else{
		log.Println("Database has been updated")
	}
}