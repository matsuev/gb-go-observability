package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AppRouter struct
type AppRouter struct {
	router *gin.Engine
	logger *zap.Logger
}

// CreateRouter function
func CreateRouter() (r *AppRouter, err error) {
	r = new(AppRouter)
	r.router = gin.Default()

	if r.logger, err = zap.NewProduction(); err != nil {
		return
	}

	err = r.initRoutes()
	return
}

// initRoutes function
func (r *AppRouter) initRoutes() (err error) {
	r.router.GET("/", r.indexHandler)
	return
}

// indexHandler function
func (r *AppRouter) indexHandler(c *gin.Context) {
	defer func() {
		if err := r.logger.Sync(); err != nil {
			log.Println(err)
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
	})

	r.logger.Info("fetch URL",
		zap.String("url", c.Request.URL.RawPath),
		zap.Duration("backoff", time.Second),
	)
}

func (r *AppRouter) Run() (err error) {
	return r.router.Run()
}
