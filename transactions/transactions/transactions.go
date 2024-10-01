package transactions

import (
	"errors"
	"time"
)

type Transaction struct {
	Amount int
	Time   time.Time
}

type TransactionsProvider interface {
	ListAllTransactions() ([]Transaction, error)
	InsertTransaction(Transaction) error
}

func CalculateBalance(provider TransactionsProvider, startTime time.Time, endTime time.Time) (int, error) {
	transactions, err := provider.ListAllTransactions()
	if err != nil {
		return 0, err
	}
	balance := 0
	for _, n := range transactions {
		if n.Time.After(startTime) && n.Time.Before(endTime) {
			balance += n.Amount
		}
	}

	return balance, nil
}

func CreateTransaction(provider TransactionsProvider, amount int) error {
	sum, _ := CalculateBalance(provider, time.Date(1, 1, 2010, 0, 0, 0, 0, time.UTC), time.Now())
	if sum+amount < 0 {
		return errors.New(("money not enaugh"))
	}
	newTransaction := Transaction{
		Amount: amount,
		Time:   time.Now(),
	}
	return provider.InsertTransaction(newTransaction)
}

// Необходимо обеспечить возможность создавать новые транзакции
//
// создаем функцию `CreateTransaction(provider TransactionProvider, amount int) error`
//
// Добавляем в TransactionProvider новый метод InsertTransaction(Transaction) error
//
// Релизуем его в inmemory provider'е
