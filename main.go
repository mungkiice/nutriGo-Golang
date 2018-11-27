package main

import "github.com/mungkiice/goNutri/route"

func main() {
	//cmd.Execute()
	route.NewRouter().Run(":8000")
}
