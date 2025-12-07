package web

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	. "github.com/mungkiice/nutriGo-Golang/database"
	. "github.com/mungkiice/nutriGo-Golang/model"
)

type loginRequest struct {
	Email string `form:"email" json:"email" binding:"required,email"`
	Pass  string `form:"password" json:"password" binding:"required"`
}

//func exists(
//	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
//	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
//) bool {
//	if err := DB.Table(param).Select("email").Where("email = ?", field.String()).Scan(&User{}).Error;
//		err == gorm.ErrRecordNotFound{
//		return false
//	}else if err != nil{
//		return false
//	}
//	return true
//}

func DoLogin(c *gin.Context) {
	var user User
	var req loginRequest
	session := sessions.Default(c)

	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//	if err := v.RegisterValidation("thereis", exists); err != nil {
	//		log.Fatal(err)
	//	}
	//}

	if err := c.ShouldBind(&req); err != nil {
		switch c.GetHeader("Content-Type") {
		case "application/json":
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		case "application/xml":
			c.XML(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		default:
			//err = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypeBind)
			//if err != nil {
			//	c.Redirect(http.StatusBadRequest, "/register")
			//}
			//c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			//	"errors" : err.Error(),
			//})
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			//c.Set("errors", err.Error())
			//c.Redirect(http.StatusMovedPermanently, "/login")
		}
		return
	}

	if err := DB.First(&user, "email = ?", req.Email).Error; err == gorm.ErrRecordNotFound {
		log.Println("user not found")
	} else if err != nil {
		log.Fatal(err)
	}
	if err := user.VerifyPassword(req.Pass); err != nil {
		log.Println("password incorrect")
	}
	session.Set("userID", user.ID)
	if err := session.Save(); err != nil {
		log.Fatal(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}
