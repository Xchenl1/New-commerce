package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId    uint `gorm:"not null"`
	ProductId uint
	BossId    uint
	AddressId uint
	Num       int
	OrderNum  uint64
	Type      uint //1：未支付 2：已支付
	Money     float64
}

func (user *Order) TableName() string {
	return "Order"
}
