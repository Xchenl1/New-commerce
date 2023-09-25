package model

import "gorm.io/gorm"

type ProductImg struct {
	gorm.Model
	ProductId uint `gorm:"not null"`
	ImgPath   string
}

func (user *ProductImg) TableName() string {
	return "ProductImg"
}
