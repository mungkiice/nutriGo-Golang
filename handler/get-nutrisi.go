package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mungkiice/goNutri/middleware/auth"
	. "github.com/mungkiice/goNutri/model"
	"net/http"
)

func NutrisiPage(c *gin.Context){
	user := auth.User(c)

	c.HTML(http.StatusOK, "nutrisi_page", gin.H{
		"user" : user,
		"isLoggedIn" : true,
		"data": getNutrisi(user),
	})
}

func getNutrisi(user *User)map[string]interface{}{
	beratIdeal := getKondisi(user)["beratIdeal"]
	var kebutuhanBasal float64
	if user.Gender == "laki-laki" {
		kebutuhanBasal = beratIdeal.(float64) * 30
	}else{
		kebutuhanBasal = beratIdeal.(float64) * 25
	}

	//50% untuk atlit, karena aplikasi ini fokus pada atlit maka menggunakan 50%
	aktifitas := kebutuhanBasal * 50 / 100
	//koreksiUsia 5% jika usia telah melebihi 40tahun
	koreksiUsia := kebutuhanBasal * 5 / 100

	totalKaloriPerhari := kebutuhanBasal + aktifitas - koreksiUsia

	//kandungan protein 10-15 % dari total Kalori ,1 kkal = 4 gram protein
	minProtein := (totalKaloriPerhari * 10 / 100) / 4
	maxProtein := (totalKaloriPerhari * 15 / 100) / 4

	//kandungan lemak 10-25 % dari total Kalori ,1 kkal = 9 gram lemak
	minLemak := (totalKaloriPerhari * 10 / 100) / 9
	maxLemak := (totalKaloriPerhari * 25 / 100) / 9

	//kandungan karbohidrat 60-75 % dari total Kalori ,1 kkal = 4 gram karbohidrat
	minKarbo := (totalKaloriPerhari * 60 / 100) / 4
	maxKarbo := (totalKaloriPerhari * 75 / 100) / 4

	return map[string]interface{}{
		"kebutuhanBasal": fmt.Sprintf("%.2f", kebutuhanBasal),
		"totalKaloriPerhari": fmt.Sprintf("%.2f", totalKaloriPerhari),
		"minProtein": fmt.Sprintf("%.2f", minProtein),
		"maxProtein": fmt.Sprintf("%.2f", maxProtein),
		"minLemak": fmt.Sprintf("%.2f", minLemak),
		"maxLemak": fmt.Sprintf("%.2f", maxLemak),
		"minKarbo": fmt.Sprintf("%.2f", minKarbo),
		"maxKarbo": fmt.Sprintf("%.2f", maxKarbo),
	}
}
