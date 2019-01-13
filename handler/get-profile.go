package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	. "github.com/mungkiice/goNutri/database"
	. "github.com/mungkiice/goNutri/model"
)

func ProfilePage(c *gin.Context){
	var user User

	userID := c.MustGet("userID").(uint)
	if  userID != 0 {
		if err := DB.First(&user, userID).Error; err != nil{
			log.Fatal(err)
		}
	}
	c.HTML(http.StatusOK, "profile_page", gin.H{
		"user": user,
		"isLoggedIn": true,
	})
}

func ProfilePageWithHistory(c *gin.Context){
	var user User
	var histories []History
	var history History

	userID := c.MustGet("userID").(uint)
	if  userID != 0 {
		if err := DB.First(&user, userID).Error; err != nil{
			log.Fatal(err)
		}
	}

	DB.Model(&user).Related(&histories).Last(&history)
	DB.Preload("Pola.Makanans").Preload("Pola").Find(&history)
	c.HTML(http.StatusOK, "profile_page", gin.H{
		"user": user,
		"isLoggedIn": true,
		"history": history,
	})
}