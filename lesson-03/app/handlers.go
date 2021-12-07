package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// indexHandler function
func (a *TracingApp) indexHandler(ctx *gin.Context) {
	user, err := a.storage.GetUserByEmail(ctx, "alex.matsuev@gmail.com")
	if err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
