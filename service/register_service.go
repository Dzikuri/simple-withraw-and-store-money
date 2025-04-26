package service

import "github.com/dzikuri/simple-withdraw-and-store-money/repository"

type RegisterService interface{}

type registerService struct {
	NasabahRepository repository.NasabahRepository
}

func NewRegisterService(nasabahRepository repository.NasabahRepository) RegisterService {
	return &registerService{NasabahRepository: nasabahRepository}
}
