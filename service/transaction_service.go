package service

import (
	"context"
	"strconv"

	"github.com/dzikuri/simple-withdraw-and-store-money/repository"
	"github.com/dzikuri/simple-withdraw-and-store-money/util"
)

type TransactionService interface {
	CheckSaldo(ctx context.Context, nasabahId string) (int64, error)
	DepositMoney(ctx context.Context, nasabahId string, amount int64) (int64, error)
	WithdrawMoney(ctx context.Context, nasabahId string, amount int64) (int64, error)
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

func (t *transactionService) DepositMoney(ctx context.Context, nasabahId string, amount int64) (int64, error) {

	// NOTE: Convert string to int64
	nasabahIdInt, err := strconv.ParseInt(nasabahId, 10, 64)
	if err != nil {
		return 0, err
	}

	// NOTE: Check if nasabah exist
	_, err = t.NasabahRepository.GetNasabahByRekeningNumber(ctx, nasabahIdInt)
	if err != nil {
		return 0, err
	}

	saldo, err := t.TransactionRepository.DepositMoney(ctx, nasabahId, amount)
	if err != nil {
		return 0, err
	}
	return saldo, nil
}

func (t *transactionService) WithdrawMoney(ctx context.Context, nasabahId string, amount int64) (int64, error) {

	// NOTE: Convert string to int64
	nasabahIdInt, err := strconv.ParseInt(nasabahId, 10, 64)
	if err != nil {
		return 0, err
	}

	// NOTE: Check if nasabah exist
	dataNasabah, err := t.NasabahRepository.GetNasabahByRekeningNumber(ctx, nasabahIdInt)
	if err != nil {
		return 0, err
	}

	if (dataNasabah.TotalMoney - amount) < 0 {
		return 0, util.ErrInsufficientBalance
	}

	saldo, err := t.TransactionRepository.WithdrawMoney(ctx, nasabahId, amount)
	if err != nil {
		return 0, err
	}
	return saldo, nil
}
