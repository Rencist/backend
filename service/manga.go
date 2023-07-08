package service

import (
	"context"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"
	"tamiyochi-backend/repository"

	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type MangaService interface {
	CreateManga(ctx context.Context, mangaDTO dto.MangaCreateDTO) (entity.Manga, error)
	GetAllManga(ctx context.Context) ([]entity.Manga, error)
	DeleteManga(ctx context.Context, mangaID uuid.UUID) (error)
	UpdateManga(ctx context.Context, mangaDTO dto.MangaUpdateDTO) (error)
}

type mangaService struct {
	mangaRepository repository.MangaRepository
}

func NewMangaService(ur repository.MangaRepository) MangaService {
	return &mangaService{
		mangaRepository: ur,
	}
}

func(us *mangaService) CreateManga(ctx context.Context, mangaDTO dto.MangaCreateDTO) (entity.Manga, error) {
	manga := entity.Manga{}
	err := smapping.FillStruct(&manga, smapping.MapFields(mangaDTO))
	if err != nil {
		return manga, err
	}
	return us.mangaRepository.CreateManga(ctx, manga)
}

func(us *mangaService) GetAllManga(ctx context.Context) ([]entity.Manga, error) {
	return us.mangaRepository.GetAllManga(ctx)
}

func(us *mangaService) DeleteManga(ctx context.Context, mangaID uuid.UUID) (error) {
	return us.mangaRepository.DeleteManga(ctx, mangaID)
}

func(us *mangaService) UpdateManga(ctx context.Context, mangaDTO dto.MangaUpdateDTO) (error) {
	manga := entity.Manga{}
	err := smapping.FillStruct(&manga, smapping.MapFields(mangaDTO))
	if err != nil {
		return err
	}
	return us.mangaRepository.UpdateManga(ctx, manga)
}