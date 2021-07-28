package iugu

import (
	"bytes"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// Customer struct
type Customer struct {
	ID                     string          `json:"id,omitempty"`
	Email                  string          `json:"email"`
	Name                   string          `json:"name"`
	Notes                  string          `json:"notes"`
	CreatedAt              string          `json:"created_at,omitempty"`
	UpdatedAt              string          `json:"updated_at,omitempty"`
	CcEmails               string          `json:"cc_emails,omitempty"`
	Phone                  int32           `json:"phone"`
	PhonePrefix            int32           `json:"phone_prefix"`
	ZipCode                string          `json:"zip_code"`
	CpfCnpj                string          `json:"cpf_cnpj"`
	District               string          `json:"district"`
	State                  string          `json:"state"`
	City                   string          `json:"city"`
	Street                 string          `json:"street"`
	Number                 string          `json:"number"`
	Complement             string          `json:"complement"`
	DefaultPaymentMethodID string          `json:"default_payment_method_id,omitempty"`
	PaymentMethods         []PaymentMethod `json:"payment_methods,omitempty"`
}

// ResponseGetCustomers struct
type ResponseGetCustomers struct {
	TotalItems int32      `json:"totalItems"`
	Items      []Customer `json:"items"`
}

// GetCustomers retrieve customers
func GetCustomers() (*[]Customer, error) {
	data, err := doRequest("GET", "/v1/customers", nil)
	if err == nil {
		str := string(data)
		res := ResponseGetCustomers{}
		json.Unmarshal([]byte(str), &res)
		return &res.Items, nil
	}
	return nil, err
}

// GetCustomerByID retrieve customers
func GetCustomerByID(id string) (*Customer, error) {
	data, err := doRequest("GET", "/v1/customers/"+id, nil)
	if err == nil {
		str := string(data)
		res := Customer{}
		json.Unmarshal([]byte(str), &res)
		return &res, nil
	}
	return nil, err
}

// GetPaymentMethodsByCustomerID retrieve customers
func GetPaymentMethodsByCustomerID(id string) (*[]PaymentMethod, error) {
	data, err := doRequest("GET", "/v1/customers/"+id+"/payment_methods", nil)
	if err == nil {
		str := string(data)
		res := []PaymentMethod{}
		json.Unmarshal([]byte(str), &res)
		return &res, nil
	}
	return nil, err
}

// CreateCustomer create customers
func CreateCustomer(customer *Customer) (*Customer, error) {
	jsonValue, _ := json.Marshal(customer)
	data, err := doRequest("POST", "/v1/customers", bytes.NewBuffer(jsonValue))
	if err == nil {
		str := string(data)
		res := Customer{}
		json.Unmarshal([]byte(str), &res)
		log.Println(string(data))
		return &res, nil
	}
	return nil, err
}
