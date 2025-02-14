package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	fare "github.com/bruhyan/gober/internal/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")

	api.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api.POST("/fare", fare.GetFare)
}
