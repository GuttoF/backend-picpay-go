package models

type Transaction struct {
	ID         int64   `json:"id"`
	Value      float64 `json:"value"`
	PayerID    int64   `json:"payer_id"`
	PayeeID    int64   `json:"payee_id"`
	Status     string  `json:"status"` // pending, authorized, failed
	Authorized int64   `json:"authorized"`
}
