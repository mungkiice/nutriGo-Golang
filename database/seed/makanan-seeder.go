package seed

import (
	"log"
	"math"
	"math/rand"

	"github.com/icrowley/fake"
	. "github.com/mungkiice/nutriGo-Golang/database"
	"github.com/mungkiice/nutriGo-Golang/model"
)

func seedMakananTable(count int) {
	for i := count; i > 0; i-- {
		DB.Create(&model.Makanan{
			Nama:        fake.ProductName(),
			Lemak:       math.Round(0+rand.Float64()*100000) / 100,
			Protein:     math.Round(0+rand.Float64()*100000) / 100,
			Karbohidrat: math.Round(0+rand.Float64()*100000) / 100,
		})
	}
	log.Println("Seed Makanan Table", count, "Rows")
}
