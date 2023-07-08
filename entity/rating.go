package entity

import "github.com/google/uuid"

type Rating struct {
	ID 		uuid.UUID 	`gorm:"primary_key;not_null" json:"id"`
	Rating 	float32 	`json:"rating"`

	SeriID int    `gorm:"foreignKey" json:"seri_id"`
	Seri   *Seri `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"seri,omitempty"`

	UserID uuid.UUID `gorm:"foreignKey" json:"user_id"`
	User   *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`

	Timestamp
}