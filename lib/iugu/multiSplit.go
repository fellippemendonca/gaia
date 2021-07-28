package iugu

type Split struct {
	ID             string `json:"id"`
	SplittableID   string `json:"splittable_id"`
	SplittableType string `json:"splittable_type"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	SplitRules     []struct {
		RecipientAccountID string  `json:"recipient_account_id"`
		Cents              int     `json:"cents,omitempty"`
		Percent            float64 `json:"percent,omitempty"`
	} `json:"split_rules"`
}
