package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Profile(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "you got me!",
	})
}
