package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type CheckoutOrder struct {
	UUID       string `json:"uuid" gorm:"unique;not null;index;primary_key"`
	Code       string `json:"code" gorm:"unique;not null;index"`
	TotalPrice uint   `json:"total_price"`

	gorm.Model
}

func (c *CheckoutOrder) BeforeCreate(tx *gorm.DB) (err error) {
	c.UUID = uuid.New().String()
	return nil
}
