package iugu

import (
	"bytes"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// Item struct
type InvoiceItem struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	PriceCents  int32  `json:"price_cents"`
	Quantity    int32  `json:"quantity"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Price       string `json:"price"`
}

// Variable struct
type Variable struct {
	Id       string `json:"id"`
	Variable string `json:"variable"`
	Value    string `json:"value"`
}

// Invoice struct
type Invoice struct {
	ID              string      `json:"id"`
	DueDate         string      `json:"due_date"`
	Currency        string      `json:"currency"`
	DiscountCents   interface{} `json:"discount_cents"`
	Email           string      `json:"email"`
	ItemsTotalCents int32       `json:"items_total_cents"`
	NotificationURL interface{} `json:"notification_url"`
	ReturnURL       interface{} `json:"return_url"`
	Status          string      `json:"status"`
	TaxCents        interface{} `json:"tax_cents"`
	UpdatedAt       string      `json:"updated_at"`
	TotalCents      int32       `json:"total_cents"`
	TotalPaidCents  int         `json:"total_paid_cents"`
	PaidAt          interface{} `json:"paid_at"`
	Pix             struct {
		Qrcode     string `json:"qrcode"`
		QrcodeText string `json:"qrcode_text"`
	} `json:"pix"`
	TaxesPaidCents            int         `json:"taxes_paid_cents"`
	PaidCents                 int         `json:"paid_cents"`
	CcEmails                  interface{} `json:"cc_emails"`
	FinancialReturnDate       string      `json:"financial_return_date"`
	PayableWith               string      `json:"payable_with"`
	OverpaidCents             interface{} `json:"overpaid_cents"`
	IgnoreDueEmail            bool        `json:"ignore_due_email"`
	IgnoreCanceledEmail       bool        `json:"ignore_canceled_email"`
	AdvanceFeeCents           interface{} `json:"advance_fee_cents"`
	EarlyPaymentDiscount      bool        `json:"early_payment_discount"`
	CommissionCents           interface{} `json:"commission_cents"`
	SecureID                  string      `json:"secure_id"`
	SecureURL                 string      `json:"secure_url"`
	CustomerID                interface{} `json:"customer_id"`
	UserID                    interface{} `json:"user_id"`
	Total                     string      `json:"total"`
	TotalPaid                 string      `json:"total_paid"`
	TotalOverpaid             string      `json:"total_overpaid"`
	TaxesPaid                 string      `json:"taxes_paid"`
	FinesOnOccurrenceDay      string      `json:"fines_on_occurrence_day"`
	TotalOnOccurrenceDay      string      `json:"total_on_occurrence_day"`
	FinesOnOccurrenceDayCents int         `json:"fines_on_occurrence_day_cents"`
	TotalOnOccurrenceDayCents int         `json:"total_on_occurrence_day_cents"`
	AdvanceFee                interface{} `json:"advance_fee"`
	Paid                      string      `json:"paid"`
	Commission                string      `json:"commission"`
	Interest                  interface{} `json:"interest"`
	Discount                  interface{} `json:"discount"`
	CreatedAt                 string      `json:"created_at"`
	Refundable                interface{} `json:"refundable"`
	Installments              interface{} `json:"installments"`
	TransactionNumber         int32       `json:"transaction_number"`
	PaymentMethod             string      `json:"payment_method"`
	CreatedAtIso              string      `json:"created_at_iso"`
	UpdatedAtIso              string      `json:"updated_at_iso"`
	OccurrenceDate            string      `json:"occurrence_date"`
	FinancialReturnDates      interface{} `json:"financial_return_dates"`
	BankSlip                  struct {
		DigitableLine string `json:"digitable_line"`
		BarcodeData   string `json:"barcode_data"`
		Barcode       string `json:"barcode"`
	} `json:"bank_slip"`
	Items           []InvoiceItem `json:"items"`
	CustomVariables []interface{} `json:"custom_variables"`
	Logs            []interface{} `json:"logs"`

	Method string   `json:"method"`
	Data   CardData `json:"data"`
	Test   bool     `json:"test"`

	WorkdayDueDate bool       `json:"ensure_workday_due_date"`
	Variables      []Variable `json:"variables"`
}

// GetAllInvoices create invoices
func GetAllInvoices() {
	data, err := doRequest("GET", "/v1/invoices", nil)
	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
	} else {
		log.Println(string(data))
	}
}

// CreateInvoice create invoices
func CreateInvoice(invoice *Invoice) {
	jsonValue, _ := json.Marshal(invoice)
	data, err := doRequest("POST", "/v1/invoices", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
	} else {
		log.Println(string(data))
	}
}

// GetInvoice get invoice
func GetInvoice(invoiceID string) {
	data, err := doRequest("POST", "/v1/invoices"+invoiceID, nil)
	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
	} else {
		log.Println(string(data))
	}
}
