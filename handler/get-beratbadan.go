package handler

import (
	"github.com/gin-gonic/gin"
	. "github.com/mungkiice/goNutri/database"
	. "github.com/mungkiice/goNutri/model"
	"log"
	"math"
	"net/http"
)

func BeratBadanPage(c *gin.Context){
	var user User
	userID := c.MustGet("userID").(uint)
	if  userID != 0 {
		if err := DB.First(&user, userID).Error; err != nil{
			log.Fatal(err)
		}
	}
	result := getKondisi(user)
	c.HTML(http.StatusOK, "beratbadan_page", gin.H{
		"user": user,
		"isLoggedIn": true,
		"kondisi": result["kondisi"],
		"beratIdeal": result["beratIdeal"],
	})
}

func getKondisi(user User)map[string]interface{}{
	berat := user.BeratBadan
	tinggi := user.TinggiBadan
	imt := berat / math.Pow((tinggi / 100),2)
	var kondisi string
	switch {
	case imt < 17:
		kondisi = "Sangat Kurus"
	case imt <= 18.5:
		kondisi = "Kurus"
	case imt <= 25:
		kondisi = "Normal (Ideal)"
	case imt <= 27:
		kondisi = "Gemuk"
	default:
		kondisi = "Sangat Gemuk (Obesitas)"
	}

	beratIdeal := 0.9 * (tinggi - 100)
	return map[string]interface{}{
		"kondisi": kondisi,
		"beratIdeal": beratIdeal,
	}
}