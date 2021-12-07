package main

import (
	"lesson-03/app"
	"log"
)

func main() {
	println("lesson-03")

	TraceApp, err := app.Create()
	if err != nil {
		log.Fatalln()
	}

	log.Fatalln(TraceApp.Run())
}
