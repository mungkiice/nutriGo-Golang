package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	. "github.com/mungkiice/goNutri/model"
	. "github.com/mungkiice/goNutri/database"
	"reflect"
)

type registerRequest struct {
	Nama		string	`form:"nama" json:"nama" binding:"required"`
	Email		string	`form:"email" json:"email" binding:"required,email,unique=user"`
	Password	string	`form:"password" json:"password" binding:"required,eqfield=PasswordC,min=6"`
	PasswordC	string	`form:"password_confirmation" json:"password_confirmation" binding:"required"`
	Gender		string	`form:"gender" json:"gender" binding:"required"`
	Usia 		int		`form:"usia" json:"usia" binding:"required"`
	TinggiBadan float64	`form:"tinggi_badan" json:"tinggi_badan" binding:"required"`
	BeratBadan 	float64	`form:"berat_badan" json:"berat_badan" binding:"required"`
}

func unique(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
	) bool {
	if err := DB.Table(param).Select("email").Where("email = ?", field.String()).Scan(&User{}).Error;
		err == gorm.ErrRecordNotFound{
		return true
	}
	return false
}

func DoRegister(c *gin.Context){
	session := sessions.Default(c)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("unique", unique); err != nil {
			log.Fatal(err)
		}
	}
	var req registerRequest

	if err := c.ShouldBind(&req); err != nil{
		switch c.GetHeader("Content-Type") {
		case "application/json":
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
		case "application/xml":
			c.XML(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
		default:
			//err = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypeBind)
			//if err != nil {
			//	c.Redirect(http.StatusBadRequest, "/register")
			//}
			//c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			//	"errors" : err.Error(),
			//})
			c.Set("errors", err.Error())
			c.Redirect(http.StatusMovedPermanently, "/register")
		}
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	user := User{
		Nama: req.Nama,
		Email: req.Email,
		Password: string(hashedPassword),
		Gender: req.Gender,
		Usia: req.Usia,
		TinggiBadan: req.TinggiBadan,
		BeratBadan: req.BeratBadan,
	}

	if err := DB.Create(&user).Error; err != nil{
		log.Fatal(err)
		c.Redirect(http.StatusMovedPermanently, "/register")
	}
	session.Set("userID", user.ID)
	if err := session.Save(); err != nil {
		log.Fatal(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}
