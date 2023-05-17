package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type CheckoutOrder struct {
	UUID       string  `json:"uuid" gorm:"unique;not null;index;primary_key"`
	TotalPrice uint    `json:"total_price"`
	Orders     []Order `json:"lines" gorm:"foreignkey:order_uuid;association_foreignkey:uuid;save_associations:false"`
	Status     string  `json:"status"`

	gorm.Model
}

type CheckoutOrderResponse struct {
	UUID       string  `json:"uuid"`
	Orders     []Order `json:"order"`
	TotalPrice uint    `json:"total_price"`
	Status     string  `json:"status"`
}

type CheckoutOrderBodyRequest struct {
	Orders []Order `json:"lines,omitempty" validate:"required"`
}

func (c *CheckoutOrder) BeforeCreate(tx *gorm.DB) (err error) {
	c.UUID = uuid.New().String()
	c.Status = "new"
	return nil
}
