package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/mungkiice/nutriGo-Golang/database"
	"github.com/mungkiice/nutriGo-Golang/middleware/auth"
	. "github.com/mungkiice/nutriGo-Golang/model"
)

func ProfilePage(c *gin.Context) {
	user := auth.User(c)
	c.HTML(http.StatusOK, "profile_page", gin.H{
		"user":       user,
		"isLoggedIn": true,
	})
}

func ProfilePageWithHistory(c *gin.Context) {
	var histories []History
	var history History

	user := auth.User(c)
	DB.Model(user).Related(&histories).Last(&history)
	DB.Preload("Pola.Makanans").Preload("Pola").Find(&history)
	c.HTML(http.StatusOK, "profile_page", gin.H{
		"user":       user,
		"isLoggedIn": true,
		"history":    history,
	})
}
