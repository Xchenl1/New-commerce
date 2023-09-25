package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserId  uint   //收货人id
	Name    string //收货人名字
	Phone   string //收货人电话
	Address string //收货人
}

func (user *Address) TableName() string {
	return "address"
}
