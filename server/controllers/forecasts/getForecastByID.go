package forecasts

import (
	log "github.com/sirupsen/logrus"

	"github.com/fellippemendonca/gaia/services"
	"github.com/gin-gonic/gin"

	"github.com/fellippemendonca/gaia/db/postgresql/orm/orm_forecast"
)

// GetForecastByID get Local Database forecast by ID
func GetForecastByID(svc *services.Services) func(c *gin.Context) {
	return func(c *gin.Context) {
		forecastID := c.Params.ByName("id")

		forecast, err := orm_forecast.GetByID(svc, forecastID)

		if err == nil {
			c.JSON(200, forecast)
			return
		}
		log.Printf("The HTTP request failed with error %s\n", err)
		c.JSON(400, nil)
	}
}
