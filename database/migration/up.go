package migration

import (
	"log"

	. "github.com/mungkiice/nutriGo-Golang/database"
	"github.com/mungkiice/nutriGo-Golang/model"
)

func Up() {
	db := DB.AutoMigrate(
		&model.User{},
		&model.Makanan{},
		&model.Pola{},
		&model.Article{},
		&model.History{})
	if err := db.Error; err != nil {
		log.Println("An Error has occured", err)
	} else {
		log.Println("Database has been updated")
	}
}
