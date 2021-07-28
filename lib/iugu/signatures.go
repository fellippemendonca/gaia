package iugu

type Feat struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Subitem struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	PriceCents  int    `json:"price_cents"`
	Price       string `json:"price"`
	Total       string `json:"total"`
}

type Log struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Notes       string `json:"notes"`
	CreatedAt   string `json:"created_at"`
}

type Signatures struct {
	ID              string        `json:"id"`
	Suspended       bool          `json:"suspended"`
	PlanIdentifier  string        `json:"plan_identifier"`
	PriceCents      int           `json:"price_cents"`
	Currency        string        `json:"currency"`
	Features        []Feat        `json:"features"`
	ExpiresAt       interface{}   `json:"expires_at"`
	CreatedAt       string        `json:"created_at"`
	UpdatedAt       string        `json:"updated_at"`
	CustomerName    string        `json:"customer_name"`
	CustomerEmail   string        `json:"customer_email"`
	CycledAt        interface{}   `json:"cycled_at"`
	CreditsMin      int           `json:"credits_min"`
	CreditsCycle    interface{}   `json:"credits_cycle"`
	CustomerID      string        `json:"customer_id"`
	PlanName        string        `json:"plan_name"`
	CustomerRef     string        `json:"customer_ref"`
	PlanRef         string        `json:"plan_ref"`
	Active          bool          `json:"active"`
	InTrial         interface{}   `json:"in_trial"`
	Credits         int           `json:"credits"`
	CreditsBased    bool          `json:"credits_based"`
	RecentInvoices  interface{}   `json:"recent_invoices"`
	Subitems        []Subitem     `json:"subitems"`
	Logs            []Log         `json:"logs"`
	CustomVariables []interface{} `json:"custom_variables"`
}
