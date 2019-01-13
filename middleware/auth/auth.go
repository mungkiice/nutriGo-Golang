package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mungkiice/goNutri/database"
	"github.com/mungkiice/goNutri/model"
	"net/http"
)

func Web() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("userID")
		if userID == nil {
			//c.Set("redirectUrl", c.Request.URL.Path)
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
		}else{
			c.Set("userID", userID)
			c.Next()
		}
	}
}

func Guest() gin.HandlerFunc{
	return func(c *gin.Context) {
		s := sessions.Default(c)
		if s.Get("userID") != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/")
			c.Abort()
		}else{
			c.Next()
		}
	}
}

func User(c *gin.Context) *model.User{
	var user model.User
	userID := sessions.Default(c).Get("userID")
	if  userID != nil {
		if err := database.DB.First(&user, userID.(uint)).Error; err != nil{
			//log.Fatal(err)
			return nil
		}
	}else{
		return nil
	}
	return &user
}
