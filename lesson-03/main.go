package main

import (
	"lesson-03/app"
	"log"
)

func main() {
	println("lesson-03")

	TraceApp := app.Create()

	log.Fatalln(TraceApp.Run())
}
