package dto

import (
	"github.com/google/uuid"
)

type UserCreateDto struct {
	ID        	uuid.UUID   `gorm:"primary_key" json:"id" form:"id"`
	Nama 		string 		`json:"nama" form:"nama" binding:"required"`
	Email 		string 		`json:"email" form:"email" binding:"required"`
	NoTelp 		string 		`json:"no_telp" form:"no_telp" binding:"required"`
	Password 	string  	`json:"password" form:"password" binding:"required"`
	Alamat 		string  	`json:"alamat" form:"alamat" binding:"required"`

	KabupatenID int   	`gorm:"foreignKey" json:"kabupaten_id" form:"kabupaten_id" binding:"required"`
}

type UserUpdateDto struct {
	ID        	uuid.UUID   `gorm:"primary_key" json:"id" form:"id"`
	Nama 		string 		`json:"nama" form:"nama"`
	Email 		string 		`json:"email" form:"email"`
	NoTelp 		string 		`json:"no_telp" form:"no_telp"`
	Password 	string  	`json:"password" form:"password"`
	Alamat 		string  	`json:"alamat" form:"alamat"`

	KabupatenID int   	`gorm:"foreignKey" json:"kabupaten_id" form:"kabupaten_id" binding:"required"`
}

type UserLoginDTO struct {
	Email 		string 		`json:"email" form:"email" binding:"email"`
	Password 	string  	`json:"password" form:"password" binding:"required"`
}