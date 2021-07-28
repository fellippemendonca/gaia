package main

import (
	"github.com/fellippemendonca/gaia/server"
)

func main() {

	/*
		customer := &models.Customer{
			Email:       "fellippe.mendonca@gmail.com",
			Name:        "Fellippe Mendonca",
			Notes:       "Engineer",
			Phone:       56610277,
			PhonePrefix: 157,
			ZipCode:     "02241120",
			CpfCnpj:     "26066736000156",
			Number:      "375",
			Street:      "Tres Rios",
			City:        "Rincão",
			State:       "São Paulo",
			District:    "Bebedouro",
			Complement:  "Pé de Goiaba",
		}
		iugu.CreateCustomer(customer)
	*/
	// log.Println(iugu.GetCustomers())
	// iugu.GetCustomerByID("2502228CADC443409C6D69090864CFBA")
	/*
		paymentToken := &models.PaymentToken{
			AccountID: "E49A4EA6851D472BA6E0C8CE3FB96E09",
			Method:    "credit_card",
			Data: models.CardData{
				FirstName:         "Cotton",
				LastName:          "Abreu",
				Number:            "5555555555554444",
				VerificationValue: "1234",
				Bin:               "411111",
				Month:             10,
				Year:              2020,
			},
			Test: true,
		}
		iugu.CreatePaymentToken(paymentToken)
	*/

	/*
		paymentMethod := &models.PaymentMethod{
			Description:  "Test payment 2",
			SetAsDefault: true,
			Token:        "4f30c349-e144-4e91-86f8-4f0fbcd4e6d4",
		}
		iugu.CreatePaymentMethod("2502228CADC443409C6D69090864CFBA", paymentMethod)
	*/
	//iugu.GetCustomerPaymentMethodByID("2502228CADC443409C6D69090864CFBA")

	server.Run()
}
