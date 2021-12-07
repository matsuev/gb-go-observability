package app

import (
	"github.com/gin-gonic/gin"
)

// TraceApplication struct
type TraceApplication struct {
	router *gin.Engine
}

// Create function
func Create() (a *TraceApplication) {
	a = new(TraceApplication)

	a.router = gin.Default()
	return
}

func (a *TraceApplication) Run() (err error) {
	err = a.router.Run()
	return
}
