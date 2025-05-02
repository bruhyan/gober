package handlers

import (
	"fmt"
	"net/http"

	"github.com/bruhyan/gober/internal/models"
	"github.com/gin-gonic/gin"
)

type Fare struct {
	Price float64 `json:"price"`
}

type FareRequest struct {
	Start       models.Location `json:"start"`
	Destination models.Location `json:"destination"`
}

func GetFare(c *gin.Context) {
	var request FareRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	start := request.Start
	dest := request.Destination

	fmt.Println(start, dest)

	var mockFare = Fare{Price: 21.90}
	c.IndentedJSON(http.StatusOK, mockFare)
}
