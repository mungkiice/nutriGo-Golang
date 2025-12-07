package main

import "github.com/mungkiice/nutriGo-Golang/cmd"

func main() {
	//if err := route.NewRouter().Run(":8000"); err != nil {
	//	log.Fatal(err)
	//}

	//if err := route.Run(); err != nil {
	//	log.Fatal(err)
	//}

	cmd.Execute()
}
