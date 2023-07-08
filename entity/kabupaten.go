package entity

type Kabupaten struct {
	ID   int    `gorm:"primary_key;not_null" json:"id"`
	Nama string `json:"nama"`

	ProvinsiID int       `gorm:"foreignKey" json:"provinsi_id"`
	Provinsi   *Provinsi `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"provinsi,omitempty"`

	Timestamp
}