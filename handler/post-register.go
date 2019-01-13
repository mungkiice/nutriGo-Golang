package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	. "github.com/mungkiice/goNutri/database"
	. "github.com/mungkiice/goNutri/model"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
)

func DoRegister(c *gin.Context){
	session := sessions.Default(c)
	nama := c.PostForm("nama")
	email := c.PostForm("email")
	password := c.PostForm("password")
	//passwordConfirmation := c.PostForm("password_confirmation")
	gender := c.PostForm("gender")
	usia, _ := strconv.Atoi(c.PostForm("usia"))
	tinggiBadan, _ := strconv.ParseFloat(c.PostForm("tinggi_badan"), 64)
	beratBadan, _ := strconv.ParseFloat(c.PostForm("berat_badan"), 64)


	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	user := User{
		Nama: nama,
		Email: email,
		Password: string(hashedPassword),
		Gender: gender,
		Usia: usia,
		TinggiBadan: tinggiBadan,
		BeratBadan: beratBadan,
	}

	if err := DB.Create(&user).Error; err != nil{
		log.Fatal(err)
		c.Redirect(http.StatusMovedPermanently, "/register")
	}
	session.Set("userID", user.ID)
	if err := session.Save(); err != nil {
		log.Fatal(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}
