package fare

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Fare struct {
	Price float64 `json:"price"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type FareRequest struct {
	Start       Location `json:"start"`
	Destination Location `json:"destination"`
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
