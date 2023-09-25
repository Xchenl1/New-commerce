package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username string
	Password string
	Avatar   string //存储用户头像的路径
}

func (user *Admin) TableName() string {
	return "admin"
}
