package service

import (
	"context"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"
	"tamiyochi-backend/repository"

	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type SeriService interface {
	CreateSeri(ctx context.Context, seriDTO dto.SeriCreateDTO) (entity.Seri, error)
	GetAllSeri(ctx context.Context, pagination entity.Pagination, filter []int, search string, sort string) (dto.PaginationResponse, error)
	DeleteSeri(ctx context.Context, seriID int) (error)
	UpdateSeri(ctx context.Context, seriDTO dto.SeriUpdateDTO) (error)
	FindSeriByID(ctx context.Context, seriID int) (dto.SeriResponseDTO, error)
	UpsertRating(ctx context.Context, seriID int, rating float32, userID uuid.UUID) (error)
	FindUserRating(ctx context.Context, userID uuid.UUID) ([]dto.RatingReponse, error)
	FindUserSeriRating(ctx context.Context, userID uuid.UUID, seriID int) (dto.RatingReponse, error)
}

type seriService struct {
	seriRepository repository.SeriRepository
}

func NewSeriService(ur repository.SeriRepository) SeriService {
	return &seriService{
		seriRepository: ur,
	}
}

func(us *seriService) CreateSeri(ctx context.Context, seriDTO dto.SeriCreateDTO) (entity.Seri, error) {
	seri := entity.Seri{}
	err := smapping.FillStruct(&seri, smapping.MapFields(seriDTO))
	if err != nil {
		return seri, err
	}
	seriID, err := us.seriRepository.GetIDSeriTerakhir(ctx)
	seri.ID = seriID + 1
	if err != nil {
		return seri, err
	}
	return us.seriRepository.CreateSeri(ctx, seri)
}

func(us *seriService) GetAllSeri(ctx context.Context, pagination entity.Pagination, filter []int, search string, sort string) (dto.PaginationResponse, error) {
	return us.seriRepository.GetAllSeri(ctx, pagination, filter, search, sort)
}

func(us *seriService) DeleteSeri(ctx context.Context, seriID int) (error) {
	return us.seriRepository.DeleteSeri(ctx, seriID)
}

func(us *seriService) UpdateSeri(ctx context.Context, seriDTO dto.SeriUpdateDTO) (error) {
	seri := entity.Seri{}
	err := smapping.FillStruct(&seri, smapping.MapFields(seriDTO))
	if err != nil {
		return err
	}
	return us.seriRepository.UpdateSeri(ctx, seri)
}

func(us *seriService) FindSeriByID(ctx context.Context, seriID int) (dto.SeriResponseDTO, error) {
	return us.seriRepository.FindSeriByIDDTOResponse(ctx, seriID)
}

func(us *seriService) UpsertRating(ctx context.Context, seriID int, rating float32, userID uuid.UUID) (error) {
	checkUserRating, _ := us.seriRepository.CheckRatingUser(ctx, seriID, userID)
	
	if checkUserRating {
		return us.seriRepository.UpdateRating(ctx, seriID, rating, userID)	
	} 
	return us.seriRepository.AddRating(ctx, seriID, rating, userID)	
}

func(us *seriService) FindUserRating(ctx context.Context, userID uuid.UUID) ([]dto.RatingReponse, error) {
	return us.seriRepository.FindUserRating(ctx, userID)
}

func(us *seriService) FindUserSeriRating(ctx context.Context, userID uuid.UUID, seriID int) (dto.RatingReponse, error) {
	return us.seriRepository.FindUserSeriRating(ctx, userID, seriID)
}
