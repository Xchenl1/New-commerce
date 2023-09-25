package model

import "gorm.io/gorm"

// Carousel 轮播图
type Carousel struct {
	gorm.Model
	ImgPath   string
	ProductId uint `gorm:"not null"`
}

func (user *Carousel) TableName() string {
	return "Carousel"
}
