package dao

import (
	"E-commerce_system/model"
	"gorm.io/gorm"
)

// CreateAddress 创建地址
func CreateAddress(DB *gorm.DB, address *model.Address) error {
	err := DB.Model(model.Address{}).Create(&address).Error
	return err
}

// GetAddress 获取地址
func GetAddress(DB *gorm.DB, id int, address *[]*model.Address) error {
	err := DB.Model(model.Address{}).Where("id=?", id).Find(&address).Error
	return err
}

// GetAddressByUserid 获取地址
func GetAddressByUserid(DB *gorm.DB, addressmap map[string]interface{}, address *[]*model.Address) (error, int64) {
	err := DB.Model(model.Address{}).Where(addressmap).Find(&address).Error
	if err != nil {
		return err, 0
	}
	var count int64
	err = DB.Model(model.Address{}).Where(addressmap).Count(&count).Error
	if err != nil {
		return err, 0
	}
	return err, count
}

// UpdateAddressById 修改地址
func UpdateAddressById(DB *gorm.DB, id int, address *model.Address) error {
	err := DB.Model(model.Address{}).Where("id=?", id).Updates(&address).Error
	return err
}

// DeleteAddressById 删除地址
func DeleteAddressById(DB *gorm.DB, id uint, userid uint) error {
	err := DB.Model(model.Address{}).Where("id = ? and user_id = ?", id, userid).Delete(&model.Address{}).Error
	return err
}
