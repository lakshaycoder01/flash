package models

import "time"

type Customer struct {
	ID           int64 `gorm:"primaryKey"`
	Name         string
	Email        string
	CreatedDate  time.Time
	ModifiedDate time.Time
}

func (Customer) TableName() string {
	return "customer"
}
