package app

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

// AppRouter struct
type AppRouter struct {
	router     *gin.Engine
	listenAddr string
}

// CreateAppRouter function
func CreateAppRouter(cfg *AppConfig) (ar *AppRouter, err error) {
	ar = new(AppRouter)

	ar.router = gin.Default()
	ar.listenAddr = cfg.ListenAddr

	ar.router.Use(ar.traceMiddleware)
	return
}

// traceMiddleware function
func (r *AppRouter) traceMiddleware(ctx *gin.Context) {
	tr := otel.Tracer("TraceApp")
	c, span := tr.Start(ctx, "Router")
	defer span.End()

	span.SetAttributes(attribute.Key("RequestURI").String(ctx.Request.RequestURI))

	ctx.Request = ctx.Request.WithContext(c)
	ctx.Next()
}

// Run function
func (r *AppRouter) Run() (err error) {
	err = r.router.Run(r.listenAddr)
	return
}
