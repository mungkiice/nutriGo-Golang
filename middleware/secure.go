package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/unrolled/secure"
	"log"
)

func Secure() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := getSecureObject().Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			c.Abort()
			return
		}

		// Avoid header rewrite if response is a redirection.
		if status := c.Writer.Status(); status > 300 && status < 399 {
			c.Abort()
		}
	}
}

func getSecureObject() *secure.Secure {
	viper.SetConfigFile("./config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	return secure.New(secure.Options{
		IsDevelopment:viper.GetBool("app_debug"),
		FrameDeny:true,
	})
}
