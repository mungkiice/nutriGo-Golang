package migration

import (
	"log"

	. "github.com/mungkiice/nutriGo-Golang/database"
	"github.com/mungkiice/nutriGo-Golang/model"
)

func Down() {
	if DB.HasTable(&model.User{}) {
		DB.DropTable(&model.User{})
	}

	if DB.HasTable(&model.Article{}) {
		DB.DropTable(&model.Article{})
	}

	if DB.HasTable(&model.History{}) {
		DB.DropTable(&model.History{})
	}

	if DB.HasTable(&model.Makanan{}) {
		DB.DropTable(&model.Makanan{})
	}

	if DB.HasTable(&model.Pola{}) {
		DB.DropTable(&model.Pola{})
	}

	if DB.HasTable("pola_makan") {
		DB.DropTable("pola_makan")
	}
	log.Println("Database has been cleared")
}
