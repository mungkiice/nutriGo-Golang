package route

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/mungkiice/goNutri/handler"
)

func NewRouter() *gin.Engine{
	Router := gin.Default()
	Router.LoadHTMLGlob("./web/view/*.html")
	Router.HTMLRender = loadTemplates("./web/public/view")
	Router.Static("/public", "./web/public")
	Router.GET("/", handler.HomePage)
	return Router
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("home_page", "./web/view/home.html", "./web/view/nav.html")

	return r
}
