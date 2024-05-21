package models

type ShortProduct struct {
	Id       uint    `json:"id"`
	Name     string  `json:"name"`
	ImageUrl string  `json:"imageUrl"`
	OldPrice float32 `json:"oldPrice,omitempty"`
	Price    float32 `json:"price"`
}
