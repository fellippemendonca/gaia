package customers

import (
	log "github.com/sirupsen/logrus"

	iugu_customers "github.com/fellippemendonca/gaia/db/postgresql/orm/iugu_customers"
	"github.com/fellippemendonca/gaia/services"
	"github.com/gin-gonic/gin"

	iugu "github.com/fellippemendonca/gaia/lib/iugu"
)

// GetIuguCustomers get all customer from external partner Iugu
func GetIuguCustomers(svc *services.Services) func(c *gin.Context) {
	return func(c *gin.Context) {
		var customers *[]iugu.Customer
		customers, err := iugu.GetCustomers()
		if err == nil {
			c.JSON(200, customers)
			return
		}
		log.Printf("The HTTP request failed with error %s\n", err)
		c.JSON(400, nil)
		return
	}
}

// GetCustomers get all customers from local database
func GetCustomers(svc *services.Services) func(c *gin.Context) {
	return func(c *gin.Context) {

		var customers []*iugu.Customer
		customers, err := iugu_customers.List(svc)

		if err == nil {
			c.JSON(200, customers)
			return
		}
		log.Printf("The HTTP request failed with error %s\n", err)
		c.JSON(400, nil)
		return
	}
}
