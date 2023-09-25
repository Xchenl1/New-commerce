package dao

import (
	"E-commerce_system/model"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// ExistOrNotByUsername 判断用户是否存在
func ExistOrNotByUsername(DB *gorm.DB, username string, user *model.User) (bool, error) {
	var count int64
	err := DB.Where("username=?", username).Find(&user).Count(&count).Error
	if err != nil {
		logx.Error("用户查询出错！", err)
		return true, err
	}
	if count == 0 { //不存在
		return false, err
	}
	//用户存在
	return true, err
}

// CreateUser 创建用户
func CreateUser(DB *gorm.DB, user *model.User) error {
	return DB.Model(&model.User{}).Create(user).Error
}
