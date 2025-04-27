package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type NasabahTransactionRepository interface {
	DepositMoney(ctx context.Context, rekeningNumber string, amount int64) (int64, error)
	WithdrawMoney(ctx context.Context, rekeningNumber string, amount int64) (int64, error)
}

type nasabahTransactionRepository struct {
	db     *pgxpool.Pool
	logger zerolog.Logger
}

func NewNasabahTransactionRepository(db *pgxpool.Pool, logger zerolog.Logger) *nasabahTransactionRepository {
	return &nasabahTransactionRepository{db: db, logger: logger}
}

func (r *nasabahTransactionRepository) DepositMoney(ctx context.Context, rekeningNumber string, amount int64) (int64, error) {
	var saldo int64
	var idNasabah uuid.UUID

	// Start transaction
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer func() {
		// Make sure rollback is called if commit is not done
		if p := recover(); p != nil {
			tx.Rollback(ctx)
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback(ctx)
		}
	}()

	// Update saldo
	sqlStatementUpdate := `UPDATE nasabah SET total_money = total_money + $1 WHERE rekening_number = $2 RETURNING id, total_money`
	err = tx.QueryRow(ctx, sqlStatementUpdate, amount, rekeningNumber).Scan(&idNasabah, &saldo)
	if err != nil {
		return 0, err
	}

	// Insert into history
	sqlStatementInsertHistory := `INSERT INTO history_transaction_nasabah (nasabah_id, transaction_type, amount, description) VALUES ($1, $2, $3, $4)`
	_, err = tx.Exec(ctx, sqlStatementInsertHistory, idNasabah, "deposit", amount, "Deposit money")
	if err != nil {
		return 0, err
	}

	// Commit transaction
	err = tx.Commit(ctx)
	if err != nil {
		return 0, err
	}

	return saldo, nil
}
func (r *nasabahTransactionRepository) WithdrawMoney(ctx context.Context, rekeningNumber string, amount int64) (int64, error) {
	var saldo int64
	var idNasabah uuid.UUID

	// Start transaction
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer func() {
		// Make sure rollback is called if commit is not done
		if p := recover(); p != nil {
			tx.Rollback(ctx)
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback(ctx)
		}
	}()

	// Update saldo
	sqlStatementUpdate := `UPDATE nasabah SET total_money = total_money - $1 WHERE rekening_number = $2 RETURNING id, total_money`
	err = tx.QueryRow(ctx, sqlStatementUpdate, amount, rekeningNumber).Scan(&idNasabah, &saldo)
	if err != nil {
		return 0, err
	}

	// Insert into history
	sqlStatementInsertHistory := `INSERT INTO history_transaction_nasabah (nasabah_id, transaction_type, amount, description) VALUES ($1, $2, $3, $4)`
	_, err = tx.Exec(ctx, sqlStatementInsertHistory, idNasabah, "withdraw", amount, "Withdraw money")
	if err != nil {
		return 0, err
	}

	// Commit transaction
	err = tx.Commit(ctx)
	if err != nil {
		return 0, err
	}

	return saldo, nil
}
