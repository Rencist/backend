package repository

import (
	"context"
	"tamiyochi-backend/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PeminjamanRepository interface {
	CreatePeminjaman(ctx context.Context, peminjaman entity.Peminjaman) (entity.Peminjaman, error)
	GetAllPeminjamanUser(ctx context.Context, userID uuid.UUID) ([]entity.Peminjaman, error)
	CreatePeminjamanManga(ctx context.Context, peminjamanManga entity.PeminjamanManga) (entity.PeminjamanManga, error)
	FindPeminjamanMangaByPeminjamanID(ctx context.Context, peminjamanID uuid.UUID) ([]entity.PeminjamanManga, error)
	FindDendaByPeminjamanID(ctx context.Context, peminjamanID uuid.UUID) (entity.Denda, error)
	PaidDenda(ctx context.Context, UserID uuid.UUID, peminjamanID uuid.UUID) (error)
}

type peminjamanConnection struct {
	connection *gorm.DB
}

func NewPeminjamanRepository(db *gorm.DB) PeminjamanRepository {
	return &peminjamanConnection{
		connection: db,
	}
}

func(db *peminjamanConnection) CreatePeminjaman(ctx context.Context, peminjaman entity.Peminjaman) (entity.Peminjaman, error) {
	uc := db.connection.Create(&peminjaman)
	if uc.Error != nil {
		return entity.Peminjaman{}, uc.Error
	}
	return peminjaman, nil
}

func(db *peminjamanConnection) GetAllPeminjamanUser(ctx context.Context, userID uuid.UUID) ([]entity.Peminjaman, error) {
	var listPeminjaman []entity.Peminjaman
	tx := db.connection.Model(entity.Peminjaman{}).Where("user_id = ?", userID).Find(&listPeminjaman)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listPeminjaman, nil
}

func(db *peminjamanConnection) CreatePeminjamanManga(ctx context.Context, peminjamanManga entity.PeminjamanManga) (entity.PeminjamanManga, error) {
	uc := db.connection.Create(&peminjamanManga)
	if uc.Error != nil {
		return entity.PeminjamanManga{}, uc.Error
	}
	return peminjamanManga, nil
}

func(db *peminjamanConnection) FindPeminjamanMangaByPeminjamanID(ctx context.Context, peminjamanID uuid.UUID) ([]entity.PeminjamanManga, error) {
	var peminjamanManga []entity.PeminjamanManga
	tx := db.connection.Where("peminjaman_id = ?", peminjamanID).Find(&peminjamanManga)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return peminjamanManga, nil
}

func(db *peminjamanConnection) FindDendaByPeminjamanID(ctx context.Context, peminjamanID uuid.UUID) (entity.Denda, error) {
	var denda entity.Denda
	tx := db.connection.Where("peminjaman_id = ?", peminjamanID).Find(&denda)
	if tx.Error != nil {
		return entity.Denda{}, tx.Error
	}
	return denda, nil
}

func(db *peminjamanConnection) PaidDenda(ctx context.Context, UserID uuid.UUID, peminjamanID uuid.UUID) (error) {
	var peminjaman entity.Peminjaman
	tx := db.connection.Model(entity.Peminjaman{}).Where("user_id = ? and id = ?", UserID, peminjamanID).Take(&peminjaman)
	if tx.Error != nil {
		return tx.Error
	}
	tx = db.connection.Model(&entity.Denda{}).Where("peminjaman_id = ?", peminjamanID).Update("is_lunas", "1").Delete(&entity.Denda{})
	if tx.Error != nil {
		return tx.Error
	}
	tx = db.connection.Model(entity.Peminjaman{}).Where("id = ?", peminjamanID).Update("status_peminjaman", "Sudah Membayar Denda")
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
