package chamaObra

// PaymentMethod struct
type PaymentMethod struct {
	ID       string   `json:"id,omitempty"`
	ItemType string   `json:"itemType,omitempty"`
	Data     CardData `json:"data,omitempty"`
}
