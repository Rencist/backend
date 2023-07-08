package service

import (
	"context"
	"tamiyochi-backend/entity"
	"tamiyochi-backend/repository"
)

type KabupatenService interface {
	GetAllKabupaten(ctx context.Context) ([]entity.Kabupaten, error)
	FindKabupatenByProvinsiID(ctx context.Context, provinsiID int) ([]entity.Kabupaten, error)
}

type kabupatenService struct {
	kabupatenRepository repository.KabupatenRepository
}

func NewKabupatenService(ur repository.KabupatenRepository) KabupatenService {
	return &kabupatenService{
		kabupatenRepository: ur,
	}
}

func(us *kabupatenService) GetAllKabupaten(ctx context.Context) ([]entity.Kabupaten, error) {
	return us.kabupatenRepository.GetAllKabupaten(ctx)
}

func(us *kabupatenService) FindKabupatenByProvinsiID(ctx context.Context, provinsiID int) ([]entity.Kabupaten, error) {
	return us.kabupatenRepository.FindKabupatenByProvinsiID(ctx, provinsiID)
}