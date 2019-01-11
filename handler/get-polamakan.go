package handler

import (
	"github.com/gin-gonic/gin"
	. "github.com/mungkiice/goNutri/database"
	. "github.com/mungkiice/goNutri/model"
	"log"
	"net/http"
)

func PolaMakanPage(c *gin.Context){
	//c.HTML(http.StatusOK, "polamakan_page", nil)
	var user User
	userID := c.MustGet("userID")
	if  userID != 0 {
		if err := DB.First(&user, userID).Error; err != nil{
			log.Fatal(err)
		}
	}

	c.HTML(http.StatusOK, "polamakan_page", gin.H{
		"user" : user,
		"loggedIn" : true,
	})
}
