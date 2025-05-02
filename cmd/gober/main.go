package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bruhyan/gober/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	fmt.Println("Running Gober in " + os.Getenv("APP_ENV") + " environment")

	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run()
}
