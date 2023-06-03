package models

type AddProductRequest struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Quantity int     `gorm:"column:quantity"`
	Price    float64 `gorm:"column:price"`
}

type AddProductResponse struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Quantity int     `gorm:"column:quantity"`
	Price    float64 `gorm:"column:price"`
}

type ReadProductResponse struct {
	Products []AddProductResponse `json:"products"`
}
