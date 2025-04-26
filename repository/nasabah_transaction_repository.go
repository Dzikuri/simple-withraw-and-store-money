package repository

import "github.com/jackc/pgx/v5/pgxpool"

type NasabahTransactionRepository interface{}

type nasabahTransactionRepository struct {
	db *pgxpool.Pool
}

func NewNasabahTransactionRepository(db *pgxpool.Pool) *nasabahTransactionRepository {
	return &nasabahTransactionRepository{db: db}
}
