package iugu

type Price struct {
	CreatedAt  string `json:"created_at"`
	Currency   string `json:"currency"`
	ID         string `json:"id"`
	PlanID     string `json:"plan_id"`
	UpdatedAt  string `json:"updated_at"`
	ValueCents int    `json:"value_cents"`
}

type Feature struct {
	CreatedAt  string      `json:"created_at"`
	ID         string      `json:"id"`
	Identifier string      `json:"identifier"`
	Important  interface{} `json:"important"`
	Name       string      `json:"name"`
	PlanID     string      `json:"plan_id"`
	Position   int         `json:"position"`
	UpdatedAt  string      `json:"updated_at"`
	Value      int         `json:"value"`
}

type PlansItem struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Identifier   string  `json:"identifier"`
	Interval     int     `json:"interval"`
	IntervalType string  `json:"interval_type"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	Prices       []Price `json:"prices"`
}

type Plan struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Identifier   string      `json:"identifier"`
	Interval     int         `json:"interval"`
	IntervalType string      `json:"interval_type"`
	CreatedAt    string      `json:"created_at"`
	UpdatedAt    string      `json:"updated_at"`
	Prices       []Price     `json:"prices"`
	Features     []Feature   `json:"features"`
	TotalItems   int         `json:"totalItems"`
	Items        []PlansItem `json:"items"`
}
