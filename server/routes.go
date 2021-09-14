package server

import (
	"github.com/fellippemendonca/gaia/services"
	"github.com/gin-gonic/gin"

	forecasts "github.com/fellippemendonca/gaia/server/controllers/forecasts"
)

// InitRoutes routes
func InitRoutes(router *gin.Engine, svc *services.Services) {
	api := router.Group("api")
	{
		api.GET("loadForecasts", forecasts.GetForecastByCoord(svc))
		api.GET("forecasts", forecasts.GetForecasts(svc))
		api.GET("forecasts/:id", forecasts.GetForecastByID(svc))
	}
}
