package service

import (
	"context"

	"github.com/dzikuri/simple-withdraw-and-store-money/model"
	"github.com/dzikuri/simple-withdraw-and-store-money/repository"
	"github.com/dzikuri/simple-withdraw-and-store-money/util"
	"github.com/rs/zerolog"
)

type RegisterService interface {
	RegisterNasabah(ctx context.Context, payload *model.CreateNasabah) (string, error)
}

type registerService struct {
	NasabahRepository repository.NasabahRepository
	Logger            zerolog.Logger
}

func NewRegisterService(nasabahRepository repository.NasabahRepository, logger zerolog.Logger) RegisterService {
	return &registerService{NasabahRepository: nasabahRepository, Logger: logger}
}

func (s *registerService) RegisterNasabah(ctx context.Context, payload *model.CreateNasabah) (string, error) {

	// NOTE: Check if nasabah exist
	nasabahExist, err := s.NasabahRepository.IfNasabahExist(ctx, &model.CheckByNikOrPhoneNumber{Nik: payload.Nik, PhoneNumber: payload.PhoneNumber})
	if err != nil {
		return "", err
	}
	if nasabahExist {
		return "", util.ErrNasabahAlreadyExist
	}

	// NOTE: Create Nasabah
	rekeningNumber, err := s.NasabahRepository.Create(ctx, payload)
	if err != nil {
		return "", err
	}

	return rekeningNumber, nil

}
