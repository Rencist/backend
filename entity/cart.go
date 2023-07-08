package entity

import "github.com/google/uuid"

type Cart struct {
	ID            uuid.UUID `gorm:"primary_key;not_null" json:"id"`
	JumlahPenyewa int       `json:"jumlah_penyewa"`

	MangaID int    `gorm:"foreignKey" json:"manga_id"`
	Manga   *Manga `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"manga,omitempty"`

	UserID uuid.UUID `gorm:"foreignKey" json:"user_id"`
	User   *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`

	Timestamp
}
