package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginPage(c *gin.Context){
	c.HTML(http.StatusOK, "login_page", nil)
}
