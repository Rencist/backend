package repository

import (
	"context"
	"tamiyochi-backend/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MangaRepository interface {
	CreateManga(ctx context.Context, manga entity.Manga) (entity.Manga, error)
	GetAllManga(ctx context.Context) ([]entity.Manga, error)
	FindMangaByID(ctx context.Context, mangaID int) (entity.Manga, error)
	DeleteManga(ctx context.Context, mangaID uuid.UUID) (error)
	UpdateManga(ctx context.Context, manga entity.Manga) (error)
}

type mangaConnection struct {
	connection *gorm.DB
}

func NewMangaRepository(db *gorm.DB) MangaRepository {
	return &mangaConnection{
		connection: db,
	}
}

func(db *mangaConnection) CreateManga(ctx context.Context, manga entity.Manga) (entity.Manga, error) {
	uc := db.connection.Create(&manga)
	if uc.Error != nil {
		return entity.Manga{}, uc.Error
	}
	return manga, nil
}

func(db *mangaConnection) GetAllManga(ctx context.Context) ([]entity.Manga, error) {
	var listManga []entity.Manga
	tx := db.connection.Find(&listManga)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listManga, nil
}

func(db *mangaConnection) FindMangaByID(ctx context.Context, mangaID int) (entity.Manga, error) {
	var manga entity.Manga
	ux := db.connection.Where("id = ?", mangaID).Find(&manga)
	if ux.Error != nil {
		return manga, ux.Error
	}
	return manga, nil
}

func(db *mangaConnection) DeleteManga(ctx context.Context, mangaID uuid.UUID) (error) {
	uc := db.connection.Delete(&entity.Manga{}, &mangaID)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func(db *mangaConnection) UpdateManga(ctx context.Context, manga entity.Manga) (error) {
	uc := db.connection.Updates(&manga)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}