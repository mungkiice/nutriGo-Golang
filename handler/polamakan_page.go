package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PolaMakanPage(c *gin.Context){
	c.HTML(http.StatusOK, "polamakan_page", nil)
}
