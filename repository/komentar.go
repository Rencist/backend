package repository

import (
	"context"
	"tamiyochi-backend/common"
	"tamiyochi-backend/dto"
	"tamiyochi-backend/entity"
	"time"

	"gorm.io/gorm"
)

type KomentarRepository interface {
	GetTotalData(ctx context.Context, seriID int) (int64, error)
	CreateKomentar(ctx context.Context, komentar entity.Komentar) (entity.Komentar, error)
	FindKomentarBySeriID(ctx context.Context, seriID int, pagination entity.Pagination) (dto.PaginationResponse, error)
}

type komentarConnection struct {
	connection *gorm.DB
}

func NewKomentarRepository(db *gorm.DB) KomentarRepository {
	return &komentarConnection{
		connection: db,
	}
}

func (db *komentarConnection) GetTotalData(ctx context.Context, seriID int) (int64, error) {
	var totalData int64
	bc := db.connection.Model(&entity.Komentar{}).Where("seri_id = ?", seriID).Count(&totalData)
	if bc.Error != nil {
		return 0, bc.Error
	}
	return totalData, nil
}

func(db *komentarConnection) CreateKomentar(ctx context.Context, komentar entity.Komentar) (entity.Komentar, error) {
	uc := db.connection.Create(&komentar)
	if uc.Error != nil {
		return entity.Komentar{}, uc.Error
	}
	return komentar, nil
}

func(db *komentarConnection) FindKomentarBySeriID(ctx context.Context, seriID int, pagination entity.Pagination) (dto.PaginationResponse, error) {
	var komentar []entity.Komentar
	totalData, _ := db.GetTotalData(ctx, seriID)
	tx := db.connection.Debug().Scopes(common.Pagination(&pagination, totalData)).Where("seri_id = ?", seriID).Order("created_at desc").Find(&komentar)
	if tx.Error != nil {
		return dto.PaginationResponse{}, tx.Error
	}
	var komentarDTO dto.KomentarResponseDTO
	var komentarDTOArray []dto.KomentarResponseDTO
	for _, res := range komentar {
		resTimeCreated, err := time.Parse("2006-1-2 15:4:5", res.CreatedAt.Format("2006-1-2 15:4:5"))
		if err != nil {
			return dto.PaginationResponse{}, err
		}
		var user entity.User
		db.connection.Where("id = ?", res.UserID).First(&user)

		komentarDTO.ID = res.ID
		komentarDTO.Isi = res.Isi
		komentarDTO.UserID = res.UserID
		komentarDTO.Username = user.Nama
		komentarDTO.CreatedAt = resTimeCreated.String()[:len(resTimeCreated.String())-10]
		komentarDTOArray = append(komentarDTOArray, komentarDTO)
	}
	
	meta := dto.Meta{
		Page: pagination.Page,
		PerPage: pagination.PerPage,
		MaxPage: pagination.MaxPage,
		TotalData: totalData,
	}
	paginationResponse := dto.PaginationResponse{
		DataPerPage: komentarDTOArray,
		Meta: meta,
	}
	return paginationResponse, nil
}
