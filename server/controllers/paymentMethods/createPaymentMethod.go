package paymentMethods

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"

	iugu "github.com/fellippemendonca/gaia/lib/iugu"
	"github.com/fellippemendonca/gaia/services"
	"github.com/gin-gonic/gin"
)

type NewCustomerPaymentMethod struct {
	Customer      iugu.Customer      `json:"customer"`
	PaymentMethod iugu.PaymentMethod `json:"paymentMethod"`
}

// CreatePaymentMethod create a new customer
func CreatePaymentMethod(svc *services.Services) func(c *gin.Context) {
	return func(c *gin.Context) {

		buf := make([]byte, 1024)
		num, _ := c.Request.Body.Read(buf)
		str := string(buf[0:num])
		res := NewCustomerPaymentMethod{}
		json.Unmarshal([]byte(str), &res)

		log.Println(res)
		jsonValue, _ := json.Marshal(res)
		log.Println(string(jsonValue))

		// var customer *iugu.Customer
		// customer, err := iugu.CreateCustomer(&res)
		// if err == nil {
		// 	c.JSON(200, customer)
		// 	return
		// }
		// log.Printf("The HTTP request failed with error %s\n", err)
		// c.JSON(400, nil)
		// return
	}
}
