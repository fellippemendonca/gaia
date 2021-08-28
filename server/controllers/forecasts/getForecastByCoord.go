package forecasts

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/fellippemendonca/gaia/services"
	"github.com/gin-gonic/gin"

	openWeatherMap "github.com/fellippemendonca/gaia/lib/openWeatherMap"
)

// GetForecastByCoord get customer from external partner Iugu by ID
func GetForecastByCoord(svc *services.Services) func(c *gin.Context) {
	return func(c *gin.Context) {
		lat, error := strconv.ParseFloat(c.DefaultQuery("latitude", "0.0"), 64)
		if error != nil {
			log.Println(error)
		}
		lon, error := strconv.ParseFloat(c.DefaultQuery("longitude", "0.0"), 64)
		if error != nil {
			log.Println(error)
		}

		var responseForecast *openWeatherMap.ResponseGetForecast
		responseForecast, err := openWeatherMap.GetForecastByCoord(lat, lon)

		if err == nil {
			c.JSON(200, responseForecast)
			return
		}
		log.Printf("The HTTP request failed with error %s\n", err)
		c.JSON(400, nil)
		return
	}
}
