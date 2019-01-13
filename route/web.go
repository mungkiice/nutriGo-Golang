package route

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/mungkiice/goNutri/handler"
	"github.com/mungkiice/goNutri/middleware"
	"github.com/mungkiice/goNutri/middleware/auth"
	"github.com/spf13/viper"
	"log"
)

func Run() error{
	router := gin.Default()

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

	router.GET("/", handler.HomePage)
	router.GET("/polamakan", auth.Web(), handler.PolaMakanPage)
	router.GET("/nutrisi", auth.Web(), handler.NutrisiPage)
	router.GET("/beratbadan", auth.Web(), handler.BeratBadanPage)
	router.GET("/forminput", auth.Web(), handler.MainFormPage)
	router.GET("login", auth.Guest(), handler.LoginPage)
	router.GET("/register", auth.Guest(), handler.RegisterPage)
	router.GET("/profile/:profileID/history", auth.Web(), handler.ProfilePageWithHistory)
	router.GET("/profile/:profileID", auth.Web(), handler.ProfilePage)
	router.POST("/login", auth.Guest(), handler.DoLogin)
	router.POST("/logout", auth.Web(), handler.DoLogout)
	router.POST("/register", auth.Guest(), handler.DoRegister)
	router.POST("/forminput", auth.Web(), handler.UpdateTinggiBadan)
	//csrf := nosurf.New(router)

	//return http.ListenAndServe(":8000", csrf)
	return router.Run(":8000")
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
	r.AddFromFiles("profile_page", templatesDir + "/profile.html", templatesDir + "/nav.html")
	return r
}
