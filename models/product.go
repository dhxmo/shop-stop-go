package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Product struct {
	UUID         string `json:"uuid" gorm:"unique;not null;index;primary_key"`
	Code         string `json:"code" gorm:"unique;not null;index"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CategoryUUID string `json:"category_uuid"`

	gorm.Model
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.UUID = uuid.New().String()
	return nil
}
