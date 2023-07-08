package dto

import (
	"tamiyochi-backend/entity"
	"time"

	"github.com/google/uuid"
)

type PeminjamanCreateDTO struct {
	ID   			uuid.UUID `gorm:"primary_key;not_null" json:"id"`
	BatasPengembalian 	time.Time 	`json:"batas_pengembalian"`
	TanggalPengembalian time.Time 	`json:"tanggal_pengembalian"`
	
	StatusPeminjaman 	string 		`json:"status_peminjaman"`
	BuktiPembayaran     string 		`json:"bukti_pembayaran" binding:"required"`
	AtasNama			string 		`json:"atas_nama" binding:"required"`

	UserID 				uuid.UUID   `gorm:"foreignKey" json:"user_id"`
}

type PeminjamanReponse struct {
	IDPeminjamanManga	uuid.UUID `json:"id_peminjaman_manga"`
	IDSeri   			int	   `json:"id_seri"`
	IDPenulis			int    `json:"id_penulis"`
	IDManga				int    `json:"id_manga"`
	IDDenda				uuid.UUID    `json:"id_denda"`
	IDPeminjaman		uuid.UUID `json:"id_peminjaman"`
	Foto				string `json:"foto"`
	Judul				string `json:"judul"`
	Penulis				[]entity.Penulis `json:"penulis"`
	Volume				int    `json:"volume"`
	TanggalPeminjaman   time.Time 	`json:"tanggal_peminjaman"`
	BatasPengembalian 	time.Time 	`json:"batas_pengembalian"`
	StatusPeminjaman	string 		`json:"status_peminjaman"`
	Denda				string 		`json:"denda"`
	IsDendaLunas		string		`json:"is_denda_lunas"`
	JumlahSewa			int 		`json:"jumlah_sewa"`
}

type DendaCreateDTO struct {
	BuktiPembayaran     string 		`json:"bukti_pembayaran" binding:"required"`
	AtasNama			string 		`json:"atas_nama" binding:"required"`
	PeminjamanID		uuid.UUID	`json:"peminjaman_id" binding:"required"`
}