package models

import "time"

type Products struct {
	ID           int64 `gorm:"primaryKey"`
	Name         string
	Price        float64
	Brand        string
	Quantity     int
	CreatedDate  time.Time
	ModifiedDate time.Time
}

func (Products) TableName() string {
	return "products"
}
