package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Order struct {
	UUID      string `json:"uuid" gorm:"unique;not null;index;primary_key"`
	ProductID string `json:"product_id" gorm:"unique;not null;index"`
	OrderID   string `json:"order_id" gorm:"unique;not null;index"`
	Quantity  uint   `json:"quantity"`
	Price     uint   `json:"price"`

	gorm.Model
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.UUID = uuid.New().String()
	return nil
}
