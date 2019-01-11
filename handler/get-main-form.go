package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MainFormPage(c *gin.Context){
	c.HTML(http.StatusOK, "mainform_page", nil)
}
