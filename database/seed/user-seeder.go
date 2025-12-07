package seed

import (
	"log"
	"math/rand"

	"github.com/icrowley/fake"
	. "github.com/mungkiice/nutriGo-Golang/database"
	"github.com/mungkiice/nutriGo-Golang/model"
	"golang.org/x/crypto/bcrypt"
)

func seedUserTable(count int) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}

	for i := count; i > 0; i-- {
		DB.Create(&model.User{
			Nama:        fake.FullName(),
			Email:       fake.EmailAddress(),
			Password:    string(hashedPassword),
			Gender:      fake.Gender(),
			TinggiBadan: float64(rand.Intn(180-50) + 50),
			BeratBadan:  float64(rand.Intn(200-40) + 40),
			Usia:        rand.Intn(100),
		})
	}

	log.Println("Seed User Table", count, "Rows")
}
