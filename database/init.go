package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mungkiice/goNutri/config"
)

var DB *gorm.DB

func init() {
	config := config.Configuration
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Database.Username,
		config.Database.Password,
		config.Database.Hostname,
		config.Database.Port,
		config.Database.Name)

	db, err := gorm.Open("mysql", dbURL)

	if err != nil {
		panic(err)
	}

	//db.DB().SetMaxIdleConns(10)
	//db.DB().SetMaxOpenConns(100)
	//db.DB().SetConnMaxLifetime(time.Hour)
	DB = db
}