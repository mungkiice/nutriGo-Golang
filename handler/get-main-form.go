package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	. "github.com/mungkiice/goNutri/database"
	. "github.com/mungkiice/goNutri/model"
)

func MainFormPage(c *gin.Context){
	var user User
	userID := c.MustGet("userID").(uint)
	if  userID != 0 {
		if err := DB.First(&user, userID).Error; err != nil{
			log.Fatal(err)
		}
	}

	c.HTML(http.StatusOK, "mainform_page", gin.H{
		"user" : user,
		"isLoggedIn" : true,
	})
}
