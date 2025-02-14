package fare

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Fare struct {
	Price float64 `json:"price"`
}

func GetFare(c *gin.Context) {
	var mockFare = Fare{Price: 21.90}
	c.IndentedJSON(http.StatusOK, mockFare)
}
