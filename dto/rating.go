package dto

import "github.com/google/uuid"

type RatingCreateDTO struct {
	ID   	uuid.UUID 	`gorm:"primary_key;not_null" json:"id"`
	Rating float32 		`json:"rating" binding:"required"`
	
	SeriID int    		`gorm:"foreignKey" json:"seri_id" binding:"required"`
	UserID uuid.UUID 	`gorm:"foreignKey" json:"user_id"`
}

type RatingUpdateDTO struct {
	ID   	uuid.UUID 	`gorm:"primary_key;not_null" json:"id"`
	Rating float32 		`json:"rating" binding:"required"`
	
	SeriID int    		`gorm:"foreignKey" json:"seri_id"`
	UserID uuid.UUID 	`gorm:"foreignKey" json:"user_id"`
}

type RatingReponse struct {
	ID   	uuid.UUID 	`gorm:"primary_key;not_null" json:"id"`
	Rating float32 		`json:"rating" binding:"required"`
	
	SeriID int    		`gorm:"foreignKey" json:"seri_id"`
	UserID uuid.UUID 	`gorm:"foreignKey" json:"user_id"`
}