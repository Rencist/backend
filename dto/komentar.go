package dto

import (
	"github.com/google/uuid"
)

type KomentarCreateDTO struct {
	ID  			uuid.UUID 	`gorm:"primary_key;not_null" json:"id"`
	Isi 			string    	`json:"isi" form:"isi" binding:"required"`

	UserID 			uuid.UUID 	`gorm:"foreignKey" json:"user_id" form:"user_id"`
	SeriID 			int 		`gorm:"foreignKey" json:"seri_id" form:"seri_id" binding:"required"`
}

type KomentarUpdateDTO struct {
	ID  			uuid.UUID 	`gorm:"primary_key;not_null" json:"id"`
	Isi 			string    	`json:"isi" form:"isi"`

	UserID 			uuid.UUID 	`gorm:"foreignKey" json:"user_id" form:"user_id"`
	SeriID 			int 		`gorm:"foreignKey" json:"seri_id" form:"seri_id"`
}

type KomentarResponseDTO struct {
	ID  			uuid.UUID 	`gorm:"primary_key;not_null" json:"id"`
	Isi 			string    	`json:"isi" form:"isi"`
	Username		string    	`json:"username" form:"username"`
	CreatedAt		string    	`json:"created_at" form:"created_at"`

	UserID 			uuid.UUID 	`gorm:"foreignKey" json:"user_id" form:"user_id"`
}
