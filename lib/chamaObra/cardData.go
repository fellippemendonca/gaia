package chamaObra

// CardData struct
type CardData struct {
	Brand         string `json:"brand,omitempty"`
	HolderName    string `json:"holderName,omitempty"`
	DisplayNumber string `json:"displayNumber,omitempty"`
	FirstDigits   string `json:"firstDigits,omitempty"`
	LastDigits    string `json:"lastDigits,omitempty"`
	MaskedNumber  string `json:"maskedNumber,omitempty"`

	FirstName         string `json:"firstName,omitempty"`
	LastName          string `json:"lastName,omitempty"`
	Number            string `json:"number,omitempty"`
	VerificationValue string `json:"verificationValue,omitempty"`
	Bin               string `json:"bin,omitempty"`
	Month             int8   `json:"month,omitempty"`
	Year              int16  `json:"year,omitempty"`
}
