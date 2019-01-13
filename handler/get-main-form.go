package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mungkiice/goNutri/middleware/auth"
	"net/http"
)

func MainFormPage(c *gin.Context){
	user := auth.User(c)
	c.HTML(http.StatusOK, "mainform_page", gin.H{
		"user" : user,
		"isLoggedIn" : true,
	})
}
