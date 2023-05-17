package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type CheckoutOrder struct {
	UUID       string  `json:"uuid" gorm:"unique;not null;index;primary_key"`
	Code       string  `json:"code" gorm:"unique;not null;index"`
	TotalPrice uint    `json:"total_price"`
	Orders     []Order `json:"lines" gorm:"foreignkey:order_uuid;association_foreignkey:uuid;save_associations:false"`
	Status     string  `json:"status"`

	gorm.Model
}

type CheckoutOrderResponse struct {
	UUID       string  `json:"uuid"`
	Code       string  `json:"code"`
	Orders     []Order `json:"order"`
	TotalPrice uint    `json:"total_price"`
	Status     string  `json:"status"`
}

type CheckoutOrderBodyRequest struct {
	Orders []Order `json:"lines,omitempty" validate:"required"`
}

// type CheckoutOrderQueryRequest struct {
// 	Code   string `json:"code,omitempty" form:"code"`
// 	Status string `json:"status,omitempty" form:"active"`
// }

func (c *CheckoutOrder) BeforeCreate(tx *gorm.DB) (err error) {
	c.UUID = uuid.New().String()
	return nil
}
