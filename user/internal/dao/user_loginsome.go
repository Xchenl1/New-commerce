package dao

import (
	"E-commerce_system/model"
	"gorm.io/gorm"
)

// LookUserInfo 查看用户的信息
func LookUserInfo(DB *gorm.DB, userid int64, name string, user *model.User) error {
	err := DB.Where("id = ? and username=?", userid, name).Find(&user).Error
	return err
}

// Update UpdateNickname 用户更新昵称
func Update(DB *gorm.DB, userid int64, name string, user *model.User) error {
	err := DB.Model(model.User{}).Where("id=?", userid).Updates(&user).Error
	return err
}
func UpdateEmail(DB *gorm.DB, userid int64, email string) error {
	err := DB.Model(model.User{}).Where("id=?", userid).Update("email", email).Error
	return err
}

// GetNoticeById 根据id查notice表 1绑定邮箱。2解绑邮箱。3修改密码
func GetNoticeById(DB *gorm.DB, id int64, notice *model.Notice) error {
	err := DB.Model(model.Notice{}).Where("id=?", id).Find(&notice).Error
	return err
}
