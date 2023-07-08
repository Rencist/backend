package entity

type Penulis struct {
	ID           int    `gorm:"primary_key;not_null" json:"id"`
	NamaDepan    string `json:"nama_depan"`
	NamaBelakang string `json:"nama_belakang"`
	Peran        string `json:"peran"`

	Timestamp
}