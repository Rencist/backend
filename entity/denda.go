package entity

import "github.com/google/uuid"

type Denda struct {
	ID   			uuid.UUID 	`gorm:"primary_key;not_null" json:"id"`
	TotalDenda 		string    	`json:"total_denda"`
	IsLunas 		string    	`json:"is_lunas"`

	PeminjamanID 	uuid.UUID   `gorm:"foreignKey" json:"peminjaman_id"`
	Peminjaman   	*Peminjaman `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"peminjaman,omitempty"`

	Timestamp
}