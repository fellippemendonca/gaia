package customers

import (
	log "github.com/sirupsen/logrus"

	"github.com/fellippemendonca/gaia/services"
	"github.com/gin-gonic/gin"

	chamaObra "github.com/fellippemendonca/gaia/lib/chamaObra"
	iugu "github.com/fellippemendonca/gaia/lib/iugu"
	iugu_customers "github.com/fellippemendonca/gaia/db/postgresql/orm/iugu_customers"
)

// GetIuguCustomerByID get customer from external partner Iugu by ID
func GetIuguCustomerByID(svc *services.Services) func(c *gin.Context) {
	return func(c *gin.Context) {
		var customerID string
		customerID = c.Params.ByName("id")

		var customer *iugu.Customer
		customer, err := iugu.GetCustomerByID(customerID)

		if err == nil {
			c.JSON(200, customer)
			return
		}
		log.Printf("The HTTP request failed with error %s\n", err)
		c.JSON(400, nil)
		return
	}
}

// GetUserByID get Local Database customer by id
func GetUserByID(svc *services.Services) func(c *gin.Context) {
	return func(c *gin.Context) {
		var customerID string
		customerID = c.Params.ByName("id")

		var customer *iugu.Customer
		customer, err := iugu_customers.GetByID(svc, customerID)

		log.Println(customer)

		user := chamaObra.User{
			ID:          customer.ID,
			Email:       customer.Email,
			Name:        customer.Name,
			Phone:       customer.Phone,
			PhonePrefix: customer.PhonePrefix,
			ZipCode:     customer.ZipCode,
			Cpf:         customer.CpfCnpj,
			District:    customer.District,
			State:       customer.State,
			City:        customer.City,
			Street:      customer.Street,
			Number:      customer.Number,
			Complement:  customer.Complement,
		}

		if err == nil {
			c.JSON(200, user)
			return
		}
		log.Printf("The HTTP request failed with error %s\n", err)
		c.JSON(400, nil)
		return
	}
}
