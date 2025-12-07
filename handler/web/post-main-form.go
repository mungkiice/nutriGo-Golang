package web

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/mungkiice/nutriGo-Golang/database"
	"github.com/mungkiice/nutriGo-Golang/middleware/auth"
)

type formRequest struct {
	Tinggi float64 `form:"tinggi_badan" json:"tinggi_badan" binding:"required"`
	Berat  float64 `form:"berat_badan" json:"berat_badan" binding:"required"`
}

func UpdateTinggiBadan(c *gin.Context) {
	var req formRequest
	user := auth.User(c)

	if err := c.ShouldBind(&req); err != nil {
		switch c.GetHeader("Content-Type") {
		case "application/json":
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		case "application/xml":
			c.XML(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		default:
			//err = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypeBind)
			//if err != nil {
			//	c.Redirect(http.StatusBadRequest, "/register")
			//}
			//c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			//	"errors" : err.Error(),
			//})
			//c.Set("errors", err.Error())
			//c.Redirect(http.StatusMovedPermanently, "/register")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	if err := DB.Model(user).Updates(map[string]interface{}{
		"tinggi_badan": req.Tinggi,
		"berat_badan":  req.Berat,
	}).Error; err != nil {
		log.Fatal(err)
	}

	c.Redirect(http.StatusMovedPermanently, "/beratbadan")
}
