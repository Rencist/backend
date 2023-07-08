package service

import (
	"context"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"
	"tamiyochi-backend/repository"

	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type KomentarService interface {
	CreateKomentar(ctx context.Context, komentarDTO dto.KomentarCreateDTO) (entity.Komentar, error)
	FindKomentarBySeriID(ctx context.Context, komentarID int, pagination entity.Pagination) (dto.PaginationResponse, error)
}

type komentarService struct {
	komentarRepository repository.KomentarRepository
}

func NewKomentarService(ur repository.KomentarRepository) KomentarService {
	return &komentarService{
		komentarRepository: ur,
	}
}

func(us *komentarService) CreateKomentar(ctx context.Context, komentarDTO dto.KomentarCreateDTO) (entity.Komentar, error) {
	komentar := entity.Komentar{}
	err := smapping.FillStruct(&komentar, smapping.MapFields(komentarDTO))
	if err != nil {
		return komentar, err
	}
	komentar.ID = uuid.New()
	return us.komentarRepository.CreateKomentar(ctx, komentar)
}

func(us *komentarService) FindKomentarBySeriID(ctx context.Context, komentarID int, pagination entity.Pagination) (dto.PaginationResponse, error) {
	return us.komentarRepository.FindKomentarBySeriID(ctx, komentarID, pagination)
}