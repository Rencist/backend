package entity

import "github.com/google/uuid"

type Transaction struct {
	ID    		uuid.UUID `gorm:"primary_key;not_null" json:"id"`
	TotalPrice  int `json:"total_price"`

	UserID uuid.UUID `gorm:"foreignKey" json:"user_id"`
	User   *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`

	MovieID int `gorm:"foreignKey" json:"movie_id"`
	Movie   *Movie     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"movie,omitempty"`

	Timestamp
}