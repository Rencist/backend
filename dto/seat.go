package dto

type SeatCreateDTO struct {
	Seat int `json:"seat" binding:"required"`
}

type AvalilableSeat struct {
	Seat int `json:"seat"`
}