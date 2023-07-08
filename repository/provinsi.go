package repository

import (
	"context"
	"tamiyochi-backend/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProvinsiRepository interface {
	GetAllProvinsi(ctx context.Context) ([]entity.Provinsi, error)
	FindProvinsiByID(ctx context.Context, provinsiID uuid.UUID) (entity.Provinsi, error)
}

type provinsiConnection struct {
	connection *gorm.DB
}

func NewProvinsiRepository(db *gorm.DB) ProvinsiRepository {
	return &provinsiConnection{
		connection: db,
	}
}

func(db *provinsiConnection) GetAllProvinsi(ctx context.Context) ([]entity.Provinsi, error) {
	var listProvinsi []entity.Provinsi
	tx := db.connection.Find(&listProvinsi)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listProvinsi, nil
}

func(db *provinsiConnection) FindProvinsiByID(ctx context.Context, provinsiID uuid.UUID) (entity.Provinsi, error) {
	var provinsi entity.Provinsi
	ux := db.connection.Where("id = ?", provinsiID).Take(&provinsi)
	if ux.Error != nil {
		return provinsi, ux.Error
	}
	return provinsi, nil
}