package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	User      User    `gorm:"Favorite:UserId"` //外键 通常来说 twobi是不允许使用外键的 严重影响性能
	UserId    uint    `gorm:"not null"`
	Product   Product `gorm:"Favorite:ProductId"`
	ProductId uint    `gorm:"not null"`
	Boss      User    `gorm:"Favorite:BossId"`
	BossId    uint    `gorm:"not null"`
}

func (user *Favorite) TableName() string {
	return "Favorite"
}
