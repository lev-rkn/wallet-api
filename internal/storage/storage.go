package storage

import (
	"sync"
	"wallet-api/internal/models"
)

var transactions = struct {
	sync.RWMutex
	data map[int][]*models.Transaction
}{
	data: map[int][]*models.Transaction{
		
	},
}

var wallets = struct {
	sync.RWMutex
	data map[int]*models.Wallet
}{
	data: map[int]*models.Wallet{
		1: {
			Id:         1,
			Balance:    8761.0,
			Identified: false,
		},
		2: {
			Id:         2,
			Balance:    87061.0,
			Identified: true,
		},
	},
}
