package customers

import (
	log "github.com/sirupsen/logrus"

	"github.com/fellippemendonca/gaia/services"
	"github.com/gin-gonic/gin"

	iugu "github.com/fellippemendonca/gaia/lib/iugu"
)

// GetIuguPaymentMethodsByCustomerID get External Partner Payment methods by customer id
func GetIuguPaymentMethodsByCustomerID(svc *services.Services) func(c *gin.Context) {
	return func(c *gin.Context) {
		var customerID string
		customerID = c.Params.ByName("id")
		iugu.GetPaymentMethodsByCustomerID(customerID)
		var paymentMethods *[]iugu.PaymentMethod
		paymentMethods, err := iugu.GetPaymentMethodsByCustomerID(customerID)
		if err == nil {
			c.JSON(200, paymentMethods)
			return
		}
		log.Printf("The HTTP request failed with error %s\n", err)
		c.JSON(400, nil)
		return
	}
}
