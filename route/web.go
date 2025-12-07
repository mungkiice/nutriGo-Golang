package route

import (
	"log"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/mungkiice/nutriGo-Golang/handler/web"
	"github.com/mungkiice/nutriGo-Golang/middleware"
	"github.com/mungkiice/nutriGo-Golang/middleware/auth"
	"github.com/spf13/viper"
)

var router *gin.Engine

func Run() error {
	router = gin.Default()

	viper.SetConfigFile("./config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	store := cookie.NewStore([]byte(viper.GetString("app_key")))
	router.Use(sessions.Sessions(viper.GetString("app_name"), store))

	//router.Use(cors.Default())
	router.Use(middleware.Secure())

	router.HTMLRender = loadTemplates("./web/view")
	router.Static("/public", "./web/public")

	loadWebRoute()
	loadApiRoute()
	//csrf := nosurf.New(router)

	//return http.ListenAndServe(":8000", csrf)
	return router.Run(":8000")
}

func loadWebRoute() {
	router.GET("/", web.HomePage)
	router.GET("/polamakan", auth.Web(), web.PolaMakanPage)
	router.GET("/nutrisi", auth.Web(), web.NutrisiPage)
	router.GET("/beratbadan", auth.Web(), web.BeratBadanPage)
	router.GET("/forminput", auth.Web(), web.MainFormPage)
	router.GET("/login", auth.Guest(), web.LoginPage)
	router.GET("/register", auth.Guest(), web.RegisterPage)
	router.GET("/profile/:profileID/history", auth.Web(), web.ProfilePageWithHistory)
	router.GET("/profile/:profileID", auth.Web(), web.ProfilePage)
	router.POST("/login", auth.Guest(), web.DoLogin)
	router.POST("/logout", auth.Web(), web.DoLogout)
	router.POST("/register", auth.Guest(), web.DoRegister)
	router.POST("/forminput", auth.Web(), web.UpdateTinggiBadan)
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("home_page", templatesDir+"/home.html", templatesDir+"/nav.html")
	r.AddFromFiles("polamakan_page", templatesDir+"/polamakan.html", templatesDir+"/nav.html")
	r.AddFromFiles("nutrisi_page", templatesDir+"/nutrisi.html", templatesDir+"/nav.html")
	r.AddFromFiles("beratbadan_page", templatesDir+"/beratbadan.html", templatesDir+"/nav.html")
	r.AddFromFiles("mainform_page", templatesDir+"/main-form.html", templatesDir+"/nav.html")
	r.AddFromFiles("login_page", templatesDir+"/login.html", templatesDir+"/nav.html")
	r.AddFromFiles("register_page", templatesDir+"/register.html", templatesDir+"/nav.html")
	r.AddFromFiles("profile_page", templatesDir+"/profile.html", templatesDir+"/nav.html")
	return r
}
