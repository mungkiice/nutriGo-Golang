package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mungkiice/nutriGo-Golang/middleware/auth"
)

func HomePage(c *gin.Context) {
	user := auth.User(c)
	if user != nil {
		c.HTML(http.StatusOK, "home_page", gin.H{
			"isLoggedIn": true,
			"user":       user,
		})
		return
	}

	c.HTML(http.StatusOK, "home_page", gin.H{
		"isLoggedIn": false,
	})

}
