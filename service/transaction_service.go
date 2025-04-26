package service

import "github.com/dzikuri/simple-withdraw-and-store-money/repository"

type TransactionService interface{}

type transactionService struct {
	TransactionRepository repository.NasabahTransactionRepository
	NasabahRepository     repository.NasabahRepository
}

func NewTransactionService(transactionRepository repository.NasabahTransactionRepository, nasabahRepository repository.NasabahRepository) TransactionService {
	return &transactionService{TransactionRepository: transactionRepository, NasabahRepository: nasabahRepository}
}
