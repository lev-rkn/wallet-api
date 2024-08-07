package models

type Wallet struct {
	Id         int     `json:"id"`
	Balance    float64 `json:"balance"`
	Identified bool    `json:"identified"`
}
