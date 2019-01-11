package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WebAuth() gin.HandlerFunc {
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
