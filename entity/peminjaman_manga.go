package entity

import "github.com/google/uuid"

type PeminjamanManga struct {
	ID uuid.UUID `gorm:"primary_key;not_null" json:"id"`

	PeminjamanID uuid.UUID     `gorm:"foreignKey" json:"peminjaman_id"`
	Peminjaman   *Peminjaman `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"peminjaman,omitempty"`

	MangaID int    `gorm:"foreignKey" json:"manga_id"`
	Manga   *Manga `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"manga,omitempty"`

	Timestamp
}