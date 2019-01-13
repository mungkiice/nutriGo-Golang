package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mungkiice/goNutri/middleware/auth"
	"net/http"
)

func HomePage(c *gin.Context){
	user := auth.User(c)
	if user != nil {
		c.HTML(http.StatusOK, "home_page", gin.H{
			"isLoggedIn" : true,
			"user" : user,
		})
		return
	}

	c.HTML(http.StatusOK, "home_page", gin.H{
		"isLoggedIn" : false,
	})

}
