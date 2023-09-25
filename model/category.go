package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string
}

func (user *Category) TableName() string {
	return "Category"
}
