package main

import (
	"github.com/mungkiice/goNutri/route"
	"log"
)

func main() {
	//cmd.Execute()
	if err := route.NewRouter().Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
