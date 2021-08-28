package server

import (
	"github.com/fellippemendonca/gaia/services"
	"github.com/gin-gonic/gin"

	customers "github.com/fellippemendonca/gaia/server/controllers/customers"
	forecasts "github.com/fellippemendonca/gaia/server/controllers/forecasts"
)

// InitRoutes routes
func InitRoutes(router *gin.Engine, svc *services.Services) {
	api := router.Group("api")
	{

		api.GET("iugu/customers", customers.GetIuguCustomers(svc))
		api.GET("iugu/customers/:id", customers.GetIuguCustomerByID(svc))
		api.GET("iugu/customers/:id/paymentMethods", customers.GetIuguPaymentMethodsByCustomerID(svc))
		api.POST("customers", customers.CreateCustomer(svc))
		api.GET("customers", customers.GetCustomers(svc))
		api.GET("customers/:id", customers.GetUserByID(svc))
		api.GET("customers/:id/paymentMethods", customers.GetIuguPaymentMethodsByCustomerID(svc))
		api.GET("forecasts", forecasts.GetForecastByCoord(svc))
	}
}
