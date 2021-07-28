package iugu

import (
	"bytes"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// CardData struct
type CardData struct {
	Brand         string `json:"brand,omitempty"`
	HolderName    string `json:"holder_name,omitempty"`
	DisplayNumber string `json:"display_number,omitempty"`
	FirstDigits   string `json:"first_digits,omitempty"`
	LastDigits    string `json:"last_digits,omitempty"`
	MaskedNumber  string `json:"masked_number,omitempty"`

	FirstName         string `json:"first_name,omitempty"`
	LastName          string `json:"last_name,omitempty"`
	Number            string `json:"number,omitempty"`
	VerificationValue string `json:"verification_value,omitempty"`
	Bin               string `json:"bin,omitempty"`
	Month             int8   `json:"month,omitempty"`
	Year              int16  `json:"year,omitempty"`
}

// PaymentToken struct
type PaymentToken struct {
	AccountID string   `json:"account_id"`
	Method    string   `json:"method"`
	Data      CardData `json:"data"`
	Test      bool     `json:"test"`
}

// PaymentMethod struct
type PaymentMethod struct {
	ID           string   `json:"id,omitempty"`
	Description  string   `json:"description"`
	Token        string   `json:"token,omitempty"`
	ItemType     string   `json:"item_type,omitempty"`
	SetAsDefault bool     `json:"set_as_default,omitempty"`
	Data         CardData `json:"data,omitempty"`
}

// CreatePaymentToken create a token
func CreatePaymentToken(paymentToken *PaymentToken) {
	jsonValue, _ := json.Marshal(paymentToken)
	data, err := doRequest("POST", "/v1/payment_token", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
	} else {
		log.Println(string(data))
	}
}

// CreatePaymentMethod create customers
func CreatePaymentMethod(customerID string, paymentMethod *PaymentMethod) {
	jsonValue, _ := json.Marshal(paymentMethod)
	data, err := doRequest("POST", "/v1/customers/"+customerID+"/payment_methods", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
	} else {
		log.Println(string(data))
	}
}
