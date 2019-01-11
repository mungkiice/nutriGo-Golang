package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	. "github.com/mungkiice/goNutri/database"
	. "github.com/mungkiice/goNutri/model"
)

func HomePage(c *gin.Context){
	//c.HTML(http.StatusOK,"home_page", nil)
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID != nil {
		var user User
		if err := DB.First(&user, userID).Error; err == gorm.ErrRecordNotFound{
			session.Delete("userID")
			if err = session.Save(); err != nil{
				log.Fatal(err)
			}
		}else if err != nil{
			log.Fatal(err)
		}
		c.HTML(http.StatusOK, "home_page", gin.H{
			"isLoggedIn" : true,
			"user" : user,
		})
	}else{
		c.HTML(http.StatusOK, "home_page", gin.H{
			"isLoggedIn" : false,
		})
	}
}
