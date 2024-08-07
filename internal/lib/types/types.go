package types

import "errors"

var (
	ErrWalletNotFound   = errors.New("кошелек не найден")
	ErrExceedMaxBalance = errors.New("вы не можете превысить максимальный баланс своего кошелька")
	ErrInvalidUserId    = errors.New("невалидный идентификатор пользователя")
)

var (
	MaxBalanceUnidentified = 10000.0
	MaxBalanceIdentified   = 100000.0
)

var (
	KeyUserId = "userId"
)
