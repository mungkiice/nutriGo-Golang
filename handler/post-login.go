package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	. "github.com/mungkiice/goNutri/database"
	. "github.com/mungkiice/goNutri/model"
	"log"
	"net/http"
)

func DoLogin(c *gin.Context) {
	var user User
	session := sessions.Default(c)
	email := c.PostForm("email")
	pass := c.PostForm("password")

	if err := DB.First(&user, "email = ?", email).Error;
		err == gorm.ErrRecordNotFound {
		log.Println("user not found")
	} else if err != nil {
		log.Fatal(err)
	}
	if err := user.VerifyPassword(pass); err != nil {
		log.Println("password incorrect")
	}
	session.Set("userID", user.ID)
	if err := session.Save(); err != nil {
		log.Fatal(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}