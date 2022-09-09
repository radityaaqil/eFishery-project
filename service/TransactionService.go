package service

import (
	"efishery-project/entities"
	"efishery-project/repository"
)

type ServiceTransaction interface {
	CreateTransactionService(transaction entities.Transaction, user_id int) (entities.Transaction, error)
	InsertPaymentSlipService(user_id int, id int, path string) (entities.Transaction, error)
	GetUserTransactionsService(user_id int) ([]entities.Transaction, error)
}

type Service_Transaction struct {
	repositorytransaction repository.RepositoryTransaction
}

func NewServiceTransaction(repositorytransaction repository.RepositoryTransaction) *Service_Transaction {
	return &Service_Transaction{repositorytransaction}
}

func (s *Service_Transaction) CreateTransactionService(transaction entities.Transaction, user_id int) (entities.Transaction, error) {
	createTransaction, errcreateTransaction := s.repositorytransaction.CreateTransaction(transaction, user_id)

	if errcreateTransaction != nil {
		return createTransaction, errcreateTransaction
	}

	return createTransaction, nil
}

func (s *Service_Transaction) InsertPaymentSlipService(user_id int, id int, path string) (entities.Transaction, error) {
	createTransaction, errcreateTransaction := s.repositorytransaction.InsertPaymentSlip(user_id, id, path)

	if errcreateTransaction != nil {
		return createTransaction, errcreateTransaction
	}

	return createTransaction, nil
}

func (s *Service_Transaction) GetUserTransactionsService(user_id int) ([]entities.Transaction, error) {
	getUserTransaction, errgetUserTransaction := s.repositorytransaction.GetUserTransactions(user_id)

	if errgetUserTransaction != nil {
		return getUserTransaction, errgetUserTransaction
	}

	return getUserTransaction, nil
}
