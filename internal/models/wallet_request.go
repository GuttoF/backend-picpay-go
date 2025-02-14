package models

type BalanceRequest struct {
    WalletID string  `json:"wallet_id" validate:"required,uuid"`
    Amount   float64 `json:"amount" validate:"required,gt=0"`
}