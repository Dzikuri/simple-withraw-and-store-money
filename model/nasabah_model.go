package model

import "github.com/google/uuid"

type Nasabah struct {
	Id             uuid.UUID `json:"id"`
	RekeningNumber int64     `json:"rekening_number"`
	Name           string    `json:"name"`
	Nik            string    `json:"nik"`
	PhoneNumber    string    `json:"phone_number"`
	TotalMoney     int64     `json:"total_money"`
	CreatedAt      int64     `json:"created_at"`
	UpdatedAt      int64     `json:"updated_at"`
	DeletedAt      int64     `json:"deleted_at"`
}

func (n *Nasabah) TableName() string {
	return "nasabah"
}

type CreateNasabah struct {
	Name        string `json:"nama" validate:"required"`
	Nik         string `json:"nik" validate:"required,gt=0"`
	PhoneNumber string `json:"no_hp" validate:"required,numeric"`
}

type TransactionPayload struct {
	NasabahId string `json:"no_rekening" validate:"required,numeric"`
	Amount    int64  `json:"nominal" validate:"required,gt=0"`
}

type CheckByNikOrPhoneNumber struct {
	Nik         string `json:"nik" validate:"required,gt=0"`
	PhoneNumber string `json:"phone_number" validate:"required,numeric"`
}

type GetSaldoParameter struct {
	RekeningNumber string `param:"no_rekening" validate:"required,numeric"`
}
