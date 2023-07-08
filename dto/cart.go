package dto

import "github.com/google/uuid"

type CartCreateDTO struct {
	ID uuid.UUID `gorm:"primary_key;not_null" json:"id"`

	MangaID int `gorm:"foreignKey" json:"manga_id" binding:"required"`
	Jumlah int `json:"jumlah" binding:"required"`

	UserID uuid.UUID `gorm:"foreignKey" json:"user_id" binding:"required"`
}

type CartResponse struct {
	Cart           []Cart `json:"cart"`
	TotalPinjaman  int    `json:"total_pinjaman"`
	TotalHargaSewa int    `json:"total_harga_sewa"`
}

type Cart struct {
	ID             uuid.UUID `json:"id"`
	MangaID        int       `json:"manga_id"`
	Foto           string    `json:"foto"`
	JumlahTersedia int       `json:"jumlah_tersedia"`
	JumlahSewa     int       `json:"jumlah_sewa"`
	HargaSewa      int       `json:"harga_sewa"`
	Volume         int       `json:"volume"`
	HargaSubTotal  int       `json:"harga_sub_total"`
	JudulSeri      string    `json:"judul_seri"`
}

type JumlahMangaCart struct {
	JumlahPenyewa int `json:"jumlah_penyewa"`
	MangaID       int `json:"manga_id"`
}
