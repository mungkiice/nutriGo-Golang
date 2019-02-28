package route

import (
	"github.com/mungkiice/goNutri/handler/api"
	"github.com/mungkiice/goNutri/middleware/auth"
)

func loadApiRoute(){
	r := router.Group("/api")
	{
		r.POST("/login", api.DoLogin)
		r.GET("/profile", auth.VerifyToken(),api.Profile)
	}
}
