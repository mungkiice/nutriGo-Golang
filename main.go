package main

import (
	"github.com/mungkiice/goNutri/route"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	//cmd.Execute()
	if err := route.NewRouter().Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
