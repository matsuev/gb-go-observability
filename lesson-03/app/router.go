package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RouterEngine struct
type RouterEngine struct {
	router *gin.Engine
}

// CreateRouter function
func CreateRouter() (r *RouterEngine, err error) {
	r = new(RouterEngine)
	r.router = gin.Default()
	err = r.initRoutes()
	return
}

// initRoutes function
func (r *RouterEngine) initRoutes() (err error) {
	r.router.GET("/", r.indexHandler)
	return
}

// indexHandler function
func (r *RouterEngine) indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
	})
}
