package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterPage(c *gin.Context){
	//if err, ok := c.Get("errors"); ok{
	//	c.String(http.StatusOK, err.(string))
	//	return
	//}
	c.HTML(http.StatusOK, "register_page", nil)
}
