package dao

import (
	"E-commerce_system/model"
	"gorm.io/gorm"
)

// ExistOrNorProduct 获取商品详细信息
func ExistOrNorProduct(DB *gorm.DB, id int, product *model.Product) (bool, error) {
	var count int64
	err := DB.Model(model.Product{}).Where("id=?", id).Find(&product).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, err //不存在
	}
	return true, err
}

// CreateCart 创建购物车
func CreateCart(DB *gorm.DB, cart *model.Cart) error {
	err := DB.Model(model.Cart{}).Create(&cart).Error
	return err
}

// GetUserInfo 查看用户的信息
func GetUserInfo(DB *gorm.DB, userid int64, boss *model.User) error {
	err := DB.Model(model.User{}).Where("id = ?", userid).Find(&boss).Error
	return err
}

// ListCart 获取购物车列表
func ListCart(DB *gorm.DB, carts *[]*model.Cart) error {
	err := DB.Model(model.Cart{}).Find(&carts).Error
	return err
}

// UpdateCartNumById 更新购物车的数量
func UpdateCartNumById(DB *gorm.DB, id int, cart *model.Cart) error {
	err := DB.Model(model.Cart{}).Where("id=?", id).Updates(&cart).Error
	return err
}

// GetCartById 查询购物车
func GetCartById(DB *gorm.DB, id int, cart *model.Cart) error {
	err := DB.Model(model.Cart{}).Where("id=?", id).Find(&cart).Error
	return err
}

// DeleteCartById 删除地址
func DeleteCartById(DB *gorm.DB, id uint, userid uint) error {
	err := DB.Model(model.Cart{}).Where("id = ? and user_id = ?", id, userid).Delete(&model.Cart{}).Error
	return err
}
