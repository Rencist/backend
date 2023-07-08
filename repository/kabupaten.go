package repository

import (
	"context"
	"tamiyochi-backend/entity"

	"gorm.io/gorm"
)

type KabupatenRepository interface {
	GetAllKabupaten(ctx context.Context) ([]entity.Kabupaten, error)
	FindKabupatenByProvinsiID(ctx context.Context, provinsiID int) ([]entity.Kabupaten, error)
}

type kabupatenConnection struct {
	connection *gorm.DB
}

func NewKabupatenRepository(db *gorm.DB) KabupatenRepository {
	return &kabupatenConnection{
		connection: db,
	}
}

func(db *kabupatenConnection) GetAllKabupaten(ctx context.Context) ([]entity.Kabupaten, error) {
	var listKabupaten []entity.Kabupaten
	tx := db.connection.Find(&listKabupaten)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listKabupaten, nil
}

func(db *kabupatenConnection) FindKabupatenByProvinsiID(ctx context.Context, provinsiID int) ([]entity.Kabupaten, error) {
	var kabupaten []entity.Kabupaten
	ux := db.connection.Where("provinsi_id = ?", provinsiID).Find(&kabupaten)
	if ux.Error != nil {
		return nil, ux.Error
	}
	return kabupaten, nil
}