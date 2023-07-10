package entity

import "github.com/google/uuid"

type Seat struct {
	ID   uuid.UUID `gorm:"primary_key;not_null" json:"id"`
	Seat int `json:"seat"`

	TransactionID uuid.UUID `gorm:"foreignKey" json:"transaction_id"`
	Transaction   *Transaction     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"transaction,omitempty"`

	Timestamp
}