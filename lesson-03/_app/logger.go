package app

import (
	"go.uber.org/zap"
)

// AppLogger struct
type AppLogger struct {
	prod  *zap.Logger
	debug *zap.Logger
}

// CreateLogger function
func CreateLogger() (a *AppLogger, err error) {
	if a.debug, err = zap.NewDevelopment(); err != nil {
		return
	}

	a.prod, err = zap.NewProduction()
	return
}
