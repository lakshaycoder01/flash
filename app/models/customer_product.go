package models

import "time"

type CustomerProduct struct {
	ID           int64 `gorm:"primaryKey"`
	ProductID    int64
	CustomerID   int64
	ProductName  string
	Quantity     int
	Price        float64
	Status       string // set('CONFIRMED', 'DELIVERD', 'CANCELLED')
	CreatedDate  time.Time
	ModifiedDate time.Time

	Products *Products `gorm:"foreignKey:ProductID;references:ID"`
	Customer *Customer `gorm:"foreignKey:CustomerID;references:ID"`
}

func (CustomerProduct) TableName() string {
	return "customer_product"
}
