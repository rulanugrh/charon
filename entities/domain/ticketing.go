package domain

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Price    int    `json:"price" form:"price"`
	Capacity int    `json:"capacity" form:"capacity"`
}
