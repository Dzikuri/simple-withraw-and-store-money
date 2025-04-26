package service

import (
	"context"
	"strconv"

	"github.com/dzikuri/simple-withdraw-and-store-money/repository"
)

type TransactionService interface {
	CheckSaldo(ctx context.Context, nasabahId string) (int64, error)
}

type transactionService struct {
	TransactionRepository repository.NasabahTransactionRepository
	NasabahRepository     repository.NasabahRepository
}

func NewTransactionService(transactionRepository repository.NasabahTransactionRepository, nasabahRepository repository.NasabahRepository) TransactionService {
	return &transactionService{TransactionRepository: transactionRepository, NasabahRepository: nasabahRepository}
}

func (t *transactionService) CheckSaldo(ctx context.Context, nasabahId string) (int64, error) {
	// NOTE: Convert string to int64
	nasabahIdInt, err := strconv.ParseInt(nasabahId, 10, 64)
	if err != nil {
		return 0, err
	}

	data, err := t.NasabahRepository.GetNasabahByRekeningNumber(ctx, nasabahIdInt)
	if err != nil {
		return 0, err
	}
	return data.TotalMoney, nil
}
