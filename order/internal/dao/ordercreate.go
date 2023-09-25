package dao

import (
	"E-commerce_system/model"
	"gorm.io/gorm"
)

// CreateOrder 创建订单
func CreateOrder(DB *gorm.DB, order *model.Order) error {
	err := DB.Model(model.Order{}).Create(&order).Error
	return err
}

// GetAddressById  获取订单
func GetAddressById(DB *gorm.DB, id int, address *model.Address) error {
	err := DB.Model(model.Address{}).Where("id=?", id).Find(&address).Error
	return err
}

// GetOrderByIdAndUserId  根据用户id和订单id来获取订单
func GetOrderByIdAndUserId(DB *gorm.DB, id int, userid int, order *model.Order) error {
	err := DB.Model(model.Order{}).Where("id=? and user_id=?", id, userid).Find(&order).Error
	return err
}

// GetOrderUserExistOrNot 判断用户是否存在
func GetOrderUserExistOrNot(DB *gorm.DB, userid int) (int64, error) {
	var count int64
	err := DB.Model(model.User{}).Where("id=?", userid).Count(&count).Error
	return count, err
}

// GetProductById 获取商品
func GetProductById(DB *gorm.DB, id int, product *model.Product) error {
	err := DB.Model(model.Product{}).Where("id=?", id).Find(&product).Error
	return err
}

// DeleteOrderById 删除订单
func DeleteOrderById(DB *gorm.DB, id uint, userid int) error {
	err := DB.Model(model.Order{}).Where("id=? and user_id=?", id, userid).Delete(&model.Order{}).Error
	return err
}

// GetUserById 获取用户
func GetUserById(DB *gorm.DB, id uint, user *model.User) error {
	err := DB.Model(model.User{}).Where("id=?", id).Find(&user).Error
	return err
}

// UpdateUser 更新用户
func UpdateUser(DB *gorm.DB, id int, user *model.User) error {
	err := DB.Model(model.User{}).Where("id=?", id).Updates(&user).Error
	return err
}

// UpdateOrder 更新订单
func UpdateOrder(DB *gorm.DB, id int, ordersum int, order *model.Order) error {
	err := DB.Model(model.Order{}).Where("id =? and order_num=?", id, ordersum).Updates(&order).Error
	return err
}

// GetOrderById 获取订单
func GetOrderById(DB *gorm.DB, id int, ordersum int, order *model.Order) error {
	err := DB.Model(model.Order{}).Where("id =? and order_num=?", id, ordersum).Find(&order).Error
	return err
}
