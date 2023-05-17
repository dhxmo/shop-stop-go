package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Product struct {
	UUID         string `json:"uuid" gorm:"unique;not null;index;primary_key"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CategoryUUID string `json:"category_uuid"`
	Active       bool   `gorm:"default:true"`

	gorm.Model
}

type ProductResponse struct {
	UUID         string `json:"uuid" gorm:"unique;not null;index;primary_key"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CategoryUUID string `json:"category_uuid"`
	Active       bool   `json:"active"`
}

type ProductRequest struct {
	Name         string `json:"name" validate:"required"`
	Description  string `json:"description,omitempty"`
	CategoryUUID string `json:"category_uuid" validate:"required"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.UUID = uuid.New().String()
	return nil
}

type ProductQueryParams struct {
	Active string `json:"active,omitempty" form:"active"`
}
