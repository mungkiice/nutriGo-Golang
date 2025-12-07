package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mungkiice/nutriGo-Golang/middleware/auth"
)

func MainFormPage(c *gin.Context) {
	user := auth.User(c)
	c.HTML(http.StatusOK, "mainform_page", gin.H{
		"user":       user,
		"isLoggedIn": true,
	})
}
