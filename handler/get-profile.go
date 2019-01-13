package handler

import (
	"github.com/gin-gonic/gin"
	. "github.com/mungkiice/goNutri/database"
	"github.com/mungkiice/goNutri/middleware/auth"
	. "github.com/mungkiice/goNutri/model"
	"net/http"
)

func ProfilePage(c *gin.Context){
	user := auth.User(c)
	c.HTML(http.StatusOK, "profile_page", gin.H{
		"user": user,
		"isLoggedIn": true,
	})
}

func ProfilePageWithHistory(c *gin.Context){
	var histories []History
	var history History

	user := auth.User(c)
	DB.Model(user).Related(&histories).Last(&history)
	DB.Preload("Pola.Makanans").Preload("Pola").Find(&history)
	c.HTML(http.StatusOK, "profile_page", gin.H{
		"user": user,
		"isLoggedIn": true,
		"history": history,
	})
}