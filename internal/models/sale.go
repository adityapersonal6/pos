package models

import "time"

type Sale struct {
	ID        int       `json:"id"`
	Item      string    `json:"item"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewSale(id int, item string, quantity int, Price float64) *Sale {
	return &Sale{
		ID:        id,
		Quantity:  quantity,
		Item:      item,
		Price:     Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (s *Sale) Update(quantity int, Price float64) {
	s.Quantity = quantity
	s.Price = Price
	s.UpdatedAt = time.Now()
}
