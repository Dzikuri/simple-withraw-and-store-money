package service

import (
	"context"
	"errors"

	"github.com/dzikuri/simple-withdraw-and-store-money/model"
	"github.com/dzikuri/simple-withdraw-and-store-money/repository"
)

type RegisterService interface {
	RegisterNasabah(ctx context.Context, payload *model.CreateNasabah) (string, error)
}

type registerService struct {
	NasabahRepository repository.NasabahRepository
}

func NewRegisterService(nasabahRepository repository.NasabahRepository) RegisterService {
	return &registerService{NasabahRepository: nasabahRepository}
}

func (s *registerService) RegisterNasabah(ctx context.Context, payload *model.CreateNasabah) (string, error) {

	// NOTE: Check if payload is valid

	// NOTE: Check if nasabah exist
	nasabahExist, err := s.NasabahRepository.IfNasabahExist(ctx, &model.CheckByNikOrPhoneNumber{Nik: payload.Nik, PhoneNumber: payload.PhoneNumber})
	if err != nil {
		return "", err
	}
	if nasabahExist {
		return "", errors.New("nasabah already exist")
	}

	// NOTE: Create Nasabah
	rekeningNumber, err := s.NasabahRepository.Create(ctx, payload)
	if err != nil {
		return "", err
	}

	return rekeningNumber, nil

}
