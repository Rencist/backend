package entity

type SeriGenre struct {
	ID int `gorm:"primary_key;not_null" json:"id"`

	SeriID int   `gorm:"foreignKey" json:"seri_id"`
	Seri   *Seri `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"seri,omitempty"`

	GenreID int    `gorm:"foreignKey" json:"genre_id"`
	Genre   *Genre `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"genre,omitempty"`

	Timestamp
}