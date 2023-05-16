package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Category struct {
	UUID        string `json:"uuid" gorm:"unique;not null;index;primary_key"`
	Code        string `json:"code" gorm:"unique;not null;index"`
	Name        string `json:"name"`
	Description string `json:"description"`

	gorm.Model
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.UUID = uuid.New().String()
	return nil
}
