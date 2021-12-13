package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// indexHandler function
func (a *TracingApp) indexHandler(ctx *gin.Context) {
	span := trace.SpanFromContext(ctx.Request.Context())
	span.SetAttributes(attribute.Key("handler").String("indexHandler"))

	user, err := a.storage.GetUserByEmail(ctx, "alex.matsuev@gmail.com")
	if err != nil {
		log.Println(err)
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
