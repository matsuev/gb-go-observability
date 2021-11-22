package app

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// MainServer structure
type MainServer struct {
	Router *gin.Engine
}

func init() {
	gin.SetMode(gin.ReleaseMode)
	rand.Seed(time.Now().UnixNano())
}

// CreateMainServer function
func CreateMainServer() (ms *MainServer) {
	ms = new(MainServer)
	ms.Router = gin.Default()

	ms.initRoutes()
	return
}

// initRoutes function
func (ms *MainServer) initRoutes() {
	ms.Router.GET("/", ms.indexHandler)
	ms.Router.GET("/rand", ms.randomHandler)
}

// indexHandler function
func (ms *MainServer) indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"index": "Ok",
	})
}

// randomHandler function
func (ms *MainServer) randomHandler(c *gin.Context) {
	sleepTime := time.Millisecond * time.Duration(rand.Intn(1000))
	time.Sleep(sleepTime)
	c.JSON(http.StatusOK, gin.H{
		"random": sleepTime,
	})
}
