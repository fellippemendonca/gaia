package customers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	chamaObra "github.com/fellippemendonca/gaia/lib/chamaObra"
	iugu "github.com/fellippemendonca/gaia/lib/iugu"
	utils "github.com/fellippemendonca/gaia/lib/utils"
	iugu_customers "github.com/fellippemendonca/gaia/db/postgresql/orm/iugu_customers"
	"github.com/fellippemendonca/gaia/services"
	"github.com/gin-gonic/gin"
)

// CreateCustomer create a new customer
func CreateCustomer(svc *services.Services) func(c *gin.Context) {
	return func(c *gin.Context) {

		str, _ := utils.ParseBody(c.Request)
		user := chamaObra.User{}
		json.Unmarshal([]byte(str), &user)

		customerForm := iugu.Customer{
			Email:       user.Email,
			Name:        user.Name,
			Notes:       "",
			Phone:       user.Phone,
			PhonePrefix: user.PhonePrefix,
			ZipCode:     user.ZipCode,
			CpfCnpj:     user.Cpf,
			District:    user.District,
			State:       user.State,
			City:        user.City,
			Street:      user.Street,
			Number:      user.Number,
			Complement:  user.Complement,
		}

		var customer *iugu.Customer
		customer, err := iugu.CreateCustomer(&customerForm)
		if err == nil {
			log.Printf("Customer ID: %s, IuguID: %s created successfully at Iugo!", user.ID, customer.ID)
		}

		_, err = iugu_customers.Create(svc, user.ID, customer)
		if err == nil {
			log.Printf("Customer ID: %s, IuguID: %s created successfully at DB!", user.ID, customer.ID)
			c.JSON(http.StatusCreated, nil)
			return
		}
		log.Printf("The HTTP request failed with error %s\n", err)
		c.JSON(http.StatusExpectationFailed, nil)
		return
	}
}
