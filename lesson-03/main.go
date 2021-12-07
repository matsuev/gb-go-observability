package main

import (
	"lesson-03/app"
	"log"
)

func main() {
	AppInstance, err := app.Create()
	if err != nil {
		log.Printf("msg: %s\n", err)
	}

	if err = AppInstance.Run(); err != nil {
		log.Printf("msg: %s\n", err)
	}
}
