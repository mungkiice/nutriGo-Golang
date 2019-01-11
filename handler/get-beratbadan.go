package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BeratBadanPage(c *gin.Context){
	c.HTML(http.StatusOK, "beratbadan_page", nil)
}
