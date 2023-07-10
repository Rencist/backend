package dto

import (
	"github.com/google/uuid"
)

type UserCreateDto struct {
	ID        	uuid.UUID   `gorm:"primary_key" json:"id" form:"id"`
	Name 		string  	`json:"name" form:"name" binding:"required"`
	Username 	string  	`json:"username" form:"username" binding:"required"`
	Email 		string 		`json:"email" binding:"email" form:"email" binding:"required"`
	Age 		string		`json:"age" form:"age" binding:"required"`
	Balance 	string      `json:"balance" form:"balance"`
	Password 	string  	`json:"password" form:"password" binding:"required"`
}

type UserUpdateDto struct {
	ID        	uuid.UUID   `gorm:"primary_key" json:"id" form:"id"`
	Name 		string  	`json:"name" form:"name"`
	Username 	string  	`json:"username" form:"username"`
	Email 		string 		`json:"email" binding:"email" form:"email"`
	Age 		string		`json:"age" form:"age"`
	Balance 	string		`json:"balance" form:"balance"`
	Password 	string  	`json:"password" form:"password"`
}

type UserLoginDTO struct {
	Email 		string 		`json:"email" form:"email" binding:"email"`
	Password 	string  	`json:"password" form:"password" binding:"required"`
}

type BalanceDTO struct {
	Balance 	string			`json:"balance" form:"balance"`
}
