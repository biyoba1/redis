package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}
