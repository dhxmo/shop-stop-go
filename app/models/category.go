package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Category struct {
	UUID        string `json:"uuid" gorm:"unique;not null;index;primary_key"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	Active      bool   `json:"active" gorm:"not null;default:true"`

	gorm.Model
}

type CategoryResponse struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

type CategoryRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Active      bool   `json:"active"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.UUID = uuid.New().String()
	return nil
}
