package handler

import (
	"github.com/gin-gonic/gin"
	. "github.com/mungkiice/goNutri/database"
	"github.com/mungkiice/goNutri/middleware/auth"
	"log"
	"net/http"
	"strconv"
)

func UpdateTinggiBadan(c *gin.Context){
	user := auth.User(c)
	tinggi, _ := strconv.Atoi(c.PostForm("tinggi_badan"))
	berat, _ := strconv.Atoi(c.PostForm("berat_badan"))

	if err := DB.Model(user).Updates(map[string]interface{}{
		"tinggi_badan" : tinggi,
		"berat_badan" : berat,

	}).Error; err != nil {
		log.Fatal(err)
	}

	c.Redirect(http.StatusMovedPermanently, "/beratbadan")
}