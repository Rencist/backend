package dto

type Meta struct {
	Page      int   `json:"page"`
	PerPage   int   `json:"per_page"`
	MaxPage   int   `json:"max_page"`
	TotalData int64 `json:"total_data"`
}

type PaginationResponse struct {
	DataPerPage any  `json:"data_per_page"`
	Meta        Meta `json:"meta"`
}