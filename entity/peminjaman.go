package entity

import (
	"time"

	"github.com/google/uuid"
)

func (Peminjaman) TableName() string {
    return "peminjamans"
}

type Peminjaman struct {
	ID   				uuid.UUID 	`gorm:"primary_key;not_null" json:"id"`
	BatasPengembalian 	time.Time 	`json:"batas_pengembalian"`
	TanggalPengembalian time.Time 	`json:"tanggal_pengembalian"`
	
	StatusPeminjaman 	string 		`json:"status_peminjaman"`
	BuktiPembayaran     string 		`json:"bukti_pembayaran"`
	AtasNama			string 		`json:"atas_nama"`

	UserID 				uuid.UUID   `gorm:"foreignKey" json:"user_id"`
	User   				*User 		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`

	Timestamp
}