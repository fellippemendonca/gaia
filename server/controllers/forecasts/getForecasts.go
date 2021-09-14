package forecasts

import (
	log "github.com/sirupsen/logrus"

	"github.com/fellippemendonca/gaia/db/postgresql/orm/orm_forecast"
	"github.com/fellippemendonca/gaia/services"
	"github.com/gin-gonic/gin"
)

// GetCustomers get all Forecasts from local database
func GetForecasts(svc *services.Services) func(c *gin.Context) {
	return func(c *gin.Context) {

		var forecasts []*orm_forecast.Forecast
		forecasts, err := orm_forecast.List(svc)

		if err == nil {
			c.JSON(200, forecasts)
			return
		}
		log.Printf("The HTTP request failed with error %s\n", err)
		c.JSON(400, nil)
	}
}
