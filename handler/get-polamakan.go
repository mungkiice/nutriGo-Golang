package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	. "github.com/mungkiice/goNutri/database"
	. "github.com/mungkiice/goNutri/model"
	"log"
	"net/http"
)

func PolaMakanPage(c *gin.Context){
	var user User
	var pola Pola
	userID := c.MustGet("userID").(uint)
	if  userID != 0 {
		if err := DB.First(&user, userID).Error; err != nil{
			log.Fatal(err)
		}
	}
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
			UserID: userID,
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
