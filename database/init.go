package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
)

var DB *gorm.DB

func init() {
	viper.SetConfigFile("./config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.hostname"),
		viper.GetInt("database.port"),
		viper.GetString("database.name"))

	db, err := gorm.Open("mysql", dbURL)

	if err != nil {
		panic(err)
	}

	//db.DB().SetMaxIdleConns(10)
	//db.DB().SetMaxOpenConns(100)
	//db.DB().SetConnMaxLifetime(time.Hour)
	DB = db
}