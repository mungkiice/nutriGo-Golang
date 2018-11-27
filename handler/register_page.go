package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterPage(c *gin.Context){
	c.HTML(http.StatusOK, "register_page", nil)
}
