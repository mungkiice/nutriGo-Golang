package web

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mungkiice/nutriGo-Golang/middleware/auth"
	. "github.com/mungkiice/nutriGo-Golang/model"
)

func BeratBadanPage(c *gin.Context) {
	user := auth.User(c)
	result := getKondisi(user)
	c.HTML(http.StatusOK, "beratbadan_page", gin.H{
		"user":       user,
		"isLoggedIn": true,
		"kondisi":    result["kondisi"],
		"beratIdeal": result["beratIdeal"],
	})
}

func getKondisi(user *User) map[string]interface{} {
	berat := user.BeratBadan
	tinggi := user.TinggiBadan
	imt := berat / math.Pow((tinggi/100), 2)
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
		"kondisi":    kondisi,
		"beratIdeal": beratIdeal,
	}
}
