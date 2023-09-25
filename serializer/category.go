package serializer

import "E-commerce_system/model"

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
	CreateAt     int64  `json:"createAt"`
}

func BuildCategory(item *model.Category) Category {
	return Category{
		Id:           int(item.ID),
		CategoryName: item.CategoryName,
		CreateAt:     item.CreatedAt.Unix(),
	}
}

func BuildCategorys(items []*model.Category) []Category {
	var category []Category
	for _, v := range items {
		product := BuildCategory(v)
		category = append(category, product)
	}
	return category
}
