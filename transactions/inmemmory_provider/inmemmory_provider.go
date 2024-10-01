package inmemmory_provider

import "codewars/transactions/transactions"

type InmemmoryProvider struct {
	Transactions []transactions.Transaction
}

func (t InmemmoryProvider) ListAllTransactions() ([]transactions.Transaction, error) {
	return t.Transactions, nil
}

func (t *InmemmoryProvider) InsertTransaction(transaction transactions.Transaction) error {
	t.Transactions = append(t.Transactions, transaction)
	return nil
}