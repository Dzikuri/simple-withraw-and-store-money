package repository

import (
	"context"

	"github.com/dzikuri/simple-withdraw-and-store-money/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NasabahRepository interface {
	Create(ctx context.Context, payload *model.CreateNasabah) (string, error)
	GetNasabahById(ctx context.Context, id string) (*model.Nasabah, error)
	GetNasabahByRekeningNumber(ctx context.Context, rekeningNumber int64) (*model.Nasabah, error)
	CheckNasabahExist(ctx context.Context, payload *model.CheckByNikOrPhoneNumber) (bool, error)
}

type nasabahRepository struct {
	db *pgxpool.Pool
}

func NewNasabahRepository(db *pgxpool.Pool) *nasabahRepository {
	return &nasabahRepository{db: db}
}

func (r *nasabahRepository) Create(ctx context.Context, payload *model.CreateNasabah) (string, error) {
	return "", nil
}

func (r *nasabahRepository) GetNasabahById(ctx context.Context, id string) (*model.Nasabah, error) {
	return nil, nil
}

func (r *nasabahRepository) GetNasabahByRekeningNumber(ctx context.Context, rekeningNumber int64) (*model.Nasabah, error) {
	return nil, nil
}

func (r *nasabahRepository) CheckNasabahExist(ctx context.Context, payload *model.CheckByNikOrPhoneNumber) (bool, error) {
	return false, nil
}
