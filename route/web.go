package route

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/mungkiice/goNutri/handler"
	"github.com/mungkiice/goNutri/middleware"
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
	router.GET("/polamakan", middleware.WebAuth(), handler.PolaMakanPage)
	router.GET("/nutrisi", middleware.WebAuth(), handler.NutrisiPage)
	router.GET("/beratbadan", middleware.WebAuth(), handler.BeratBadanPage)
	router.GET("/forminput", middleware.WebAuth(), handler.MainFormPage)
	router.GET("login", middleware.Guest(), handler.LoginPage)
	router.GET("/register", middleware.Guest(), handler.RegisterPage)
	router.GET("/profile/:profileID/history", middleware.WebAuth(), handler.ProfilePageWithHistory)
	router.GET("/profile/:profileID", middleware.WebAuth(), handler.ProfilePage)
	router.POST("/login", middleware.Guest(), handler.DoLogin)
	router.POST("/logout", middleware.WebAuth(), handler.DoLogout)
	router.POST("/register", middleware.Guest(), handler.DoRegister)
	router.POST("/forminput", middleware.WebAuth(), handler.UpdateTinggiBadan)
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
