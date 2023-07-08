package entity

type Genre struct {
	ID   int    `gorm:"primary_key;not_null" json:"id"`
	Nama string `json:"nama"`

	Timestamp
}