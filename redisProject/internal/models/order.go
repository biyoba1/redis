package models

import (
	"errors"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	PersonID uint `gorm:"foreignkey:PersonID;references:ID"`
	Person   Person
	Products []Product `gorm:"many2many:order_products"`
}

type UpdateOrderInput struct {
	PersonID uint   `json:"person_id"`
	Products []uint `json:"products"`
}

func (i UpdateOrderInput) Validate() error {
	if i.PersonID == 0 {
		return errors.New("person_id is required")
	}
	if len(i.Products) == 0 {
		return errors.New("products is required")
	}
	return nil
}
