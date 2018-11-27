package route

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/mungkiice/goNutri/handler"
)

func NewRouter() *gin.Engine{
	Router := gin.Default()
	Router.HTMLRender = loadTemplates("./web/view")
	Router.Static("/public", "./web/public")


	Router.GET("/", handler.HomePage)
	Router.GET("/polamakan", handler.PolaMakanPage)
	Router.GET("/nutrisi", handler.NutrisiPage)
	Router.GET("/beratbadan", handler.BeratBadanPage)
	Router.GET("/formInput", handler.MainFormPage)
	Router.GET("login", handler.LoginPage)
	Router.GET("/register", handler.RegisterPage)
	return Router
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("home_page", templatesDir + "/home.html", templatesDir + "/nav.html")
	r.AddFromFiles("polamakan_page", templatesDir + "/polamakan.html", templatesDir + "/nav.html")
	r.AddFromFiles("nutrisi_page", templatesDir + "/nutrisi.html", templatesDir + "/nav.html")
	r.AddFromFiles("beratbadan_page", templatesDir + "/beratbadan.html", templatesDir + "/nav.html")
	r.AddFromFiles("mainform_page", templatesDir + "/main-form.html", templatesDir + "/nav.html")
	r.AddFromFiles("login_page", templatesDir + "/login.html", templatesDir + "/nav.html")
	r.AddFromFiles("register_page", templatesDir + "/register.html", templatesDir + "/nav.html")
	return r
}
