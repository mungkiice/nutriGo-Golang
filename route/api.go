package route

import (
	"github.com/mungkiice/nutriGo-Golang/handler/api"
	"github.com/mungkiice/nutriGo-Golang/middleware/auth"
)

func loadApiRoute() {
	r := router.Group("/api")
	{
		r.POST("/login", api.DoLogin)
		r.GET("/profile", auth.VerifyToken(), api.Profile)
	}
}
