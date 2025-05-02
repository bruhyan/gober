package services

import (
	"github.com/bruhyan/gober/internal/models"
)

// curl 'http://127.0.0.1:8081/route/v1/driving/103.77351357810227,1.2960562347566502;103.869557,1.370414?steps=true'
func GetRoute(start models.Location, dest models.Location) models.Location {
	return models.Location{Lat: 1, Lng: 2}
}
