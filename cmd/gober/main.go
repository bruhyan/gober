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
		// not fatal as will err in docker-compose even though .env is loaded
		log.Println(".env file not found — assuming env vars are set externally")
	}

	fmt.Println("Running Gober in " + os.Getenv("APP_ENV") + " environment")

	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run()
}
