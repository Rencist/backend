package entity

import (
	"tamiyochi-backend/helpers"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        	uuid.UUID   `gorm:"primary_key;not_null" json:"id"`
	Name 		string  	`json:"name"`
	Username 	string  	`json:"username"`
	Email 		string 		`json:"email" binding:"email"`
	Age 		int			`json:"age"`
	Password 	string  	`json:"password"`
	
	Timestamp
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}