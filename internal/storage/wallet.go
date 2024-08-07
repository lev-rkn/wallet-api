package storage

import (
	"time"
	"wallet-api/internal/lib/types"
	"wallet-api/internal/models"

	"github.com/google/uuid"
)

func FoundWallet(userId int) (*models.Wallet, error) {
	// блокируем только операции на запись в хранилище, пока мы читаем (может даже много раз)
	wallets.RLock()
	defer wallets.RUnlock()

	wallet, found := wallets.data[userId]
	if !found {
		return nil, types.ErrWalletNotFound
	}

	return wallet, nil
}

func Deposit(userId int, transaction *models.Transaction) (*models.Wallet, error) {
	// блокируем любые операции чтения и записи с кошельками и транзакциями
	wallets.Lock()
	transactions.Lock()
	defer wallets.Unlock()
	defer transactions.Unlock()

	// ищем кошелек пользователя
	wallet, found := wallets.data[userId]
	if !found {
		return nil, types.ErrWalletNotFound
	}

	futureBalance := wallet.Balance + transaction.Amount
	// проверяем, не превысит ли он пороговую сумму баланса кошелька
	var maxBalance float64
	if wallet.Identified {
		maxBalance = types.MaxBalanceIdentified
	} else {
		maxBalance = types.MaxBalanceUnidentified
	}
	if futureBalance > maxBalance {
		return nil, types.ErrExceedMaxBalance
	}

	// добавляем остальные данные в транзакцию
	transaction.Id = uuid.New().String()
	transaction.WalletID = wallet.Id
	transaction.Timestamp = time.Now()

	// мутируем хранилища
	wallets.data[userId].Balance = futureBalance
	transactions.data[userId] = append(transactions.data[userId], transaction)

	return wallet, nil
}

func GetMonthHistory(userId int,
) (totalTransactions int, totalAmountOfTransactions float64) {
	// блокируем запись транзакций, пока мы читаем историю
	transactions.RLock()
	defer transactions.RUnlock()

	userTransactions := transactions.data[userId]

	currentMonth := time.Now().Month()

	for _, transaction := range userTransactions {
		if transaction.Timestamp.Month() == currentMonth {
			totalTransactions++
			totalAmountOfTransactions += transaction.Amount
		}
	}

	return
}
