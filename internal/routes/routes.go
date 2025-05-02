package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bruhyan/gober/internal/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")

	api.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// user
	api.POST("/user", handlers.CreateUser)

	// fare
	api.POST("/fare", handlers.GetFare)

}
