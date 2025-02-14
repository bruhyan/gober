package main

import (
	"github.com/bruhyan/gober/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run()
}
