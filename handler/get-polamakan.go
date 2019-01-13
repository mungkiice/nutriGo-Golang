package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	. "github.com/mungkiice/goNutri/database"
	"github.com/mungkiice/goNutri/middleware/auth"
	. "github.com/mungkiice/goNutri/model"
	"log"
	"net/http"
)

func PolaMakanPage(c *gin.Context){
	var pola Pola
	user := auth.User(c)
	nutrisi := getNutrisi(user)
	err := DB.Preload("Makanans").
		Where("lemak_total BETWEEN ? AND ? AND protein_total BETWEEN ? AND ? AND karbohidrat_total BETWEEN ? AND ?",
		nutrisi["minLemak"],
		nutrisi["maxLemak"],
		nutrisi["minProtein"],
		nutrisi["maxProtein"],
		nutrisi["minKarbo"],
		nutrisi["maxKarbo"]).
		First(&pola).Error

	if err == nil || err != gorm.ErrRecordNotFound{
		if err = DB.Create(&History{
			UserID: user.ID,
			PolaID: pola.ID,
			TinggiBadan: user.TinggiBadan,
			BeratBadan: user.BeratBadan,
		}).Error; err != nil{
			log.Fatal(err)
		}
	}

	c.HTML(http.StatusOK, "polamakan_page", gin.H{
		"user" : user,
		"isLoggedIn" : true,
		"pola": pola,
	})
}
