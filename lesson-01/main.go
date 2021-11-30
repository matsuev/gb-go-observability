package main

import (
	"lesson-01/app"
	"log"
)

func main() {
	appServer, err := app.Create()
	if err != nil {
		log.Fatalln(err)
	}

	appServer.Run()
}
