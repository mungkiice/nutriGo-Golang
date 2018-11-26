package seed

import (
	"github.com/icrowley/fake"
	"github.com/mungkiice/goNutri/model"
	. "github.com/mungkiice/goNutri/database"
	"log"
	"math"
	"math/rand"
)

func seedMakananTable(count int){
	for i := count; i > 0; i--  {
		DB.Create(&model.Makanan{
			Nama		:	fake.ProductName(),
			Lemak		:	math.Round(0 + rand.Float64() * 100000)/100,
			Protein		:	math.Round(0 + rand.Float64() * 100000)/100,
			Karbohidrat	:	math.Round(0 + rand.Float64() * 100000)/100,
		})
	}
	log.Println("Seed Makanan Table", count, "Rows")
}
