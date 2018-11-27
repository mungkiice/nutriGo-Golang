package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NutrisiPage(c *gin.Context){
	c.HTML(http.StatusOK, "nutrisi_page", nil)
}
