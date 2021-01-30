package models

import (
	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID `json:"_id"`
	Code      string    `json:"code"`
	Price     int       `json:"price"`
	Rate      float32   `json:"rate"`
	IsSale    bool      `json:"is_sale"`
	IsPromo   bool      `json:"is_promo"`
	MainImage string    `json:"main_image"`
}
