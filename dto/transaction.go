package dto

import "github.com/google/uuid"

type TransactionCreateDTO struct {
	ID uuid.UUID `gorm:"primary_key;not_null" json:"id"`

	TotalPrice  string `json:"total_price" binding:"required"`

	UserID 	uuid.UUID `gorm:"foreignKey" json:"user_id"`
	MovieID string `gorm:"foreignKey" json:"movie_id" binding:"required"`

	Seat 	[]SeatCreateDTO `gorm:"foreignKey" json:"seat" binding:"required"`
}