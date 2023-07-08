package entity

import (
	"tamiyochi-backend/helpers"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        	uuid.UUID   `gorm:"primary_key;not_null" json:"id"`
	Nama 		string 		`json:"nama"`
	Email 		string 		`json:"email" binding:"email"`
	NoTelp 		string 		`json:"no_telp"`
	Password 	string  	`json:"password"`
	Alamat		string		`json:"alamat"`
	Peran		string		`json:"peran"`

	KabupatenID	int		`gorm:"foreignKey" json:"kabupaten_id"`
	Kabupaten   *Kabupaten  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"kabupaten,omitempty"`
	
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