package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Phone    string `gorm:"unique" binding:"required"`
	Email    string `binding:"required,email"`
	Image    string
	Name     string    `binding:"required"`
	Password string    `binding:"required"`
	Contacts []Contact `gorm:"many2many:user_contacts;"`
}
