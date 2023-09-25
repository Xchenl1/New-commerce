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
	err := DB.Model(model.User{}).Where("id=? and username=?", userid, name).Updates(&user).Error
	return err
}

// GetNoticeById 根据id查notice表 1绑定邮箱。2解绑邮箱。3修改密码
func GetNoticeById(DB *gorm.DB, id int64, notice *model.Notice) error {
	err := DB.Model(model.Notice{}).Where("id=?", id).Find(&notice).Error
	return err
}

// CreateProduct 更新商品
func CreateProduct(DB *gorm.DB, Product *model.Product) error {
	err := DB.Model(model.Product{}).Create(&Product).Error
	return err
}

// CreateProductImg 添加商品图片表
func CreateProductImg(DB *gorm.DB, ProductImg *model.ProductImg) error {
	err := DB.Model(model.ProductImg{}).Create(&ProductImg).Error
	return err
}

// ListProduct 查询商品列表
func ListProduct(DB *gorm.DB, condition map[string]interface{}, pagesize int, pagesum int, product *[]*model.Product) error {
	err := DB.Model(model.Product{}).Where(condition).
		Offset((pagesum - 1) * pagesize).
		Limit(pagesize).
		Find(&product).Error
	return err
}

// CountProduct 查询商品总数
func CountProduct(DB *gorm.DB, condition map[string]interface{}, total *int64) error {
	err := DB.Model(model.Product{}).Where(condition).Count(total).Error
	return err
}

// SelectProduct 查询商品
func SelectProduct(DB *gorm.DB, info string, product *[]*model.Product, page model.BasePage) (total int64, err error) {
	err = DB.Model(model.Product{}).Where("title LIKE ? or info LIKE ?", "%"+info+"%", "%"+info+"%").
		Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&product).Error
	if err != nil {
		return 0, err
	}
	err = DB.Model(model.Product{}).Where("title LIKE ? or info LIKE ?", "%"+info+"%", "%"+info+"%").Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, err
}

// ShowProduct 获取商品详细信息
func ShowProduct(DB *gorm.DB, id int, product *model.Product) error {
	err := DB.Model(model.Product{}).Where("id=?", id).Find(&product).Error
	return err
}

// ListProductImg 获取图片地址
func ListProductImg(DB *gorm.DB, id int, productimg *[]*model.ProductImg) error {
	err := DB.Model(model.ProductImg{}).Where("product_id=?", id).Find(&productimg).Error
	return err
}

// ListProductCategory 获取分类
func ListProductCategory(DB *gorm.DB, category *[]*model.Category) error {
	err := DB.Model(model.Category{}).Find(&category).Error
	return err
}
