package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	UUID     string `json:"uuid" gorm:"unique;not null;index;primary_key"`
	Username string `json:"username" gorm:"unique;not null;index"`
	Email    string `json:"email" gorm:"unique;not null;index"`
	Password string `json:"password"`
	RoleUUID string `json:"role_uuid"`

	gorm.Model
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New().String()
	return nil
}
