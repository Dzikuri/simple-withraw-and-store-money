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
	IfNasabahExist(ctx context.Context, payload *model.CheckByNikOrPhoneNumber) (bool, error)
}

type nasabahRepository struct {
	db *pgxpool.Pool
}

func NewNasabahRepository(db *pgxpool.Pool) *nasabahRepository {
	return &nasabahRepository{db: db}
}

func (r *nasabahRepository) Create(ctx context.Context, payload *model.CreateNasabah) (string, error) {
	var noRekening string
	err := r.db.QueryRow(ctx, `
		INSERT INTO nasabah (name, nik, phone_number)
		VALUES ($1, $2, $3)
		RETURNING rekening_number
	`, payload.Name, payload.Nik, payload.PhoneNumber).Scan(&noRekening)
	if err != nil {
		return "", err
	}
	return noRekening, nil
}

func (r *nasabahRepository) GetNasabahById(ctx context.Context, id string) (*model.Nasabah, error) {
	var data model.Nasabah
	err := r.db.QueryRow(ctx, `
        SELECT id, name, nik, phone_number, total_money
        FROM nasabah
        WHERE id = $1`, id).Scan(
		&data.Id,
		&data.Name,
		&data.Nik,
		&data.PhoneNumber,
		&data.TotalMoney)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *nasabahRepository) GetNasabahByRekeningNumber(ctx context.Context, rekeningNumber int64) (*model.Nasabah, error) {
	var data model.Nasabah
	err := r.db.QueryRow(ctx, `
        SELECT id, name, nik, phone_number, total_money WHERE rekening_number = $1`, rekeningNumber).Scan(
		&data.Id,
		&data.Name,
		&data.Nik,
		&data.PhoneNumber,
		&data.TotalMoney)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *nasabahRepository) IfNasabahExist(ctx context.Context, payload *model.CheckByNikOrPhoneNumber) (bool, error) {
	var exists bool
	err := r.db.QueryRow(ctx, `
		SELECT EXISTS (
			SELECT 1 FROM nasabah
			WHERE nik = $1 OR phone_number = $2
		)
	`, payload.Nik, payload.PhoneNumber).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
