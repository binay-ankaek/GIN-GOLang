package model

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	UserID uint   `gorm:"index"`
	Name   string `binding:"required"`
	Phone  string
}
