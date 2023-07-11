package dto

import "github.com/google/uuid"

type TransactionCreateDTO struct {
	ID uuid.UUID `gorm:"primary_key;not_null" json:"id"`

	TotalPrice  string `json:"total_price" binding:"required"`

	UserID 	uuid.UUID `gorm:"foreignKey" json:"user_id"`
	MovieID string `gorm:"foreignKey" json:"movie_id" binding:"required"`

	Seat 	[]SeatCreateDTO `gorm:"foreignKey" json:"seat" binding:"required"`
}

type TransactionResponse struct {
	ID uuid.UUID `gorm:"primary_key;not_null" json:"id"`
	MovieID int `gorm:"foreignKey" json:"movie_id" binding:"required"`
	UserName string `json:"user_name"`
	MovieName string `json:"movie_name"`
	TotalPrice  string `json:"total_price" binding:"required"`
	Seat 	[]SeatCreateDTO `gorm:"foreignKey" json:"seat" binding:"required"`
}

type DeleteTransactionDTO struct {
	TransactionID string `json:"transaction_id" binding:"required" form:"transaction_id"`
}