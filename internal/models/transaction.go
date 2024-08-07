package models

import "time"

type Transaction struct {
	Id        string    `json:"id"`
	WalletID  int       `json:"wallet_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}
