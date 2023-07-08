package service

import (
	"context"
	"tamiyochi-backend/entity"
	"tamiyochi-backend/repository"
)

type ProvinsiService interface {
	GetAllProvinsi(ctx context.Context) ([]entity.Provinsi, error)
}

type provinsiService struct {
	provinsiRepository repository.ProvinsiRepository
}

func NewProvinsiService(ur repository.ProvinsiRepository) ProvinsiService {
	return &provinsiService{
		provinsiRepository: ur,
	}
}

func(us *provinsiService) GetAllProvinsi(ctx context.Context) ([]entity.Provinsi, error) {
	return us.provinsiRepository.GetAllProvinsi(ctx)
}