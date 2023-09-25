package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserId    uint `gorm:"not null"` //用户id
	ProductId uint `gorm:"not null"` //商品id
	BossId    uint `gorm:"not null"` //商家id
	Num       uint `gorm:"not null"` //数量
	MaxNum    uint `gorm:"not null"` //限额
	Check     bool //是否支付
}

func (user *Cart) TableName() string {
	return "Cart"
}
