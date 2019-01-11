package seed

import (
	"github.com/icrowley/fake"
	. "github.com/mungkiice/goNutri/database"
	"github.com/mungkiice/goNutri/model"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
)

func seedUserTable(count int){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.DefaultCost)

	if err != nil{
		log.Fatal(err)
	}

	for i := count; i > 0; i--  {
		DB.Create(&model.User{
			Nama		:	fake.FullName(),
			Email		:	fake.EmailAddress(),
			Password	:	string(hashedPassword),
			Gender		:	fake.Gender(),
			TinggiBadan	:	rand.Intn(180-50) + 50,
			BeratBadan	:	rand.Intn(200-40) + 40,
			Usia		:	rand.Intn(100),
		})
	}

	log.Println("Seed User Table", count, "Rows")
}
