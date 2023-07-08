package entity

type PenulisSeri struct {
	ID   		int 	`gorm:"primary_key;not_null" json:"id"`

	PenulisID 	int		`gorm:"foreignKey" json:"penulis_id"`
	Penulis   	*Penulis 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"penulis,omitempty"`
	
	SeriID 		int	`gorm:"foreignKey" json:"seri_id"`
	Seri   		*Seri 		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"seri,omitempty"`

	Timestamp
}