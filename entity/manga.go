package entity

type Manga struct {
	ID             int `gorm:"primary_key;not_null" json:"id"`
	Volume         int `json:"volume"`
	JumlahTersedia int `json:"jumlah_tersedia"`
	HargaSewa      int `json:"harga_sewa"`

	SeriID int   `gorm:"foreignKey" json:"seri_id"`
	Seri   *Seri `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"seri,omitempty"`

	Timestamp
}