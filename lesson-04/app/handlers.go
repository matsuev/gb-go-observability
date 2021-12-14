package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// indexHandler function
func (a *TracingApp) indexHandler(ctx *gin.Context) {
	user := new(User)

	if err := ctx.ShouldBind(user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Printf("%#v\n", user)

	if err := a.storage.GetUserByEmail(ctx, user); err != nil {
		ctx.AbortWithError(http.StatusNoContent, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// GetAllUsers function
func (a *TracingApp) GetAllUsers(ctx *gin.Context) {
	users := new(Users)

	if err := a.storage.FindAllUsers(ctx, users); err != nil {
		ctx.AbortWithError(http.StatusNoContent, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}
