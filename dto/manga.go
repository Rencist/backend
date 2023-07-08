package dto

import "github.com/google/uuid"

type MangaCreateDTO struct {
	ID   			uuid.UUID `gorm:"primary_key;not_null" json:"id"`
	Volume 			string    `json:"volume"`
	JumlahTersedia 	string    `json:"jumlah_tersedia"`
	HargaSewa 		string    `json:"harga_sewa"`

	SeriID 			uuid.UUID   `gorm:"foreignKey" json:"seri_id"`
}

type MangaUpdateDTO struct {
	ID   			uuid.UUID `gorm:"primary_key;not_null" json:"id"`
	Volume 			string    `json:"volume"`
	JumlahTersedia 	string    `json:"jumlah_tersedia"`
	HargaSewa 		string    `json:"harga_sewa"`

	SeriID 			uuid.UUID   `gorm:"foreignKey" json:"seri_id"`
}

type MangaReponse struct {
	ID   			uuid.UUID `gorm:"primary_key;not_null" json:"id"`
	Volume 			string    `json:"volume"`
	JumlahTersedia 	string    `json:"jumlah_tersedia"`
	HargaSewa 		string    `json:"harga_sewa"`
}