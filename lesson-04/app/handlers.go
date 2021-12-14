package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// indexHandler function
func (a *TracingApp) indexHandler(ctx *gin.Context) {
	user, err := a.storage.GetUserByEmail(ctx, "alex.matsuev@gmail.com")
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (a *TracingApp) GetAllUsers(ctx *gin.Context) {
	users, err := a.storage.FindAllUsers(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusNoContent, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}
