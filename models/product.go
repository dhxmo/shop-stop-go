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

	gorm.Model
}

type ProductResponse struct {
	UUID         string `json:"uuid" gorm:"unique;not null;index;primary_key"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CategoryUUID string `json:"category_uuid"`

	gorm.Model
}

type ProductRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	CategoryUUID string `json:"category_uuid"`

	gorm.Model
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.UUID = uuid.New().String()
	return nil
}
