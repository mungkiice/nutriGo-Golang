package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	. "github.com/mungkiice/nutriGo-Golang/database"
	"github.com/mungkiice/nutriGo-Golang/middleware/auth"
	"github.com/mungkiice/nutriGo-Golang/model"
)

type loginRequest struct {
	Email string `json:"email" binding:"required"`
	Pass  string `json:"password" binding:"required"`
}

func DoLogin(c *gin.Context) {
	var req loginRequest
	var user model.User

	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := DB.First(&user, "email = ?", req.Email).Error; err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, gin.H{
			"message": "user not found",
		})
		return
	} else if err != nil {
		log.Fatal(err)
	}
	if err := user.VerifyPassword(req.Pass); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "password incorrect",
		})
		return
	}

	token, err := auth.GenerateToken(&user)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"token":  token,
		"user":   user,
	})
}
