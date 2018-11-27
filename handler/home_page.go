package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomePage(c *gin.Context){
	c.HTML(http.StatusOK,"home_page", nil)
}
