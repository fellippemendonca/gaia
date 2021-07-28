package chamaObra

// User struct
type User struct {
	ID                     string          `json:"id,omitempty"`
	Email                  string          `json:"email"`
	Name                   string          `json:"name"`
	Phone                  int32           `json:"phone"`
	PhonePrefix            int32           `json:"phonePrefix"`
	ZipCode                string          `json:"zipCode"`
	Cpf                    string          `json:"cpf"`
	District               string          `json:"district"`
	State                  string          `json:"state"`
	City                   string          `json:"city"`
	Street                 string          `json:"street"`
	Number                 string          `json:"number"`
	Complement             string          `json:"complement"`
	DefaultPaymentMethodID string          `json:"defaultPaymentMethodID,omitempty"`
	PaymentMethods         []PaymentMethod `json:"paymentMethods,omitempty"`
}
