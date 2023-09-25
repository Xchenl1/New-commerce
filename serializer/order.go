package serializer

import (
	"E-commerce_system/model"
	"fmt"
)

type Order struct {
	Id           int    `json:"id"`
	OrderNum     uint64 `json:"orderNum"`
	CreatedAt    int64  `json:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt"`
	UserId       uint   `json:"userId"`
	ProductId    uint   `json:"productId"`
	BossId       uint   `json:"bossId"`
	Num          int    `json:"num"`
	AddressName  string `json:"addressName"`
	AddressPhone string `json:"addressPhone"`
	Address      string `json:"address"`
	Type         uint   `json:"type"`
	ProductName  string `json:"productName"`
	ImgPath      string `json:"imgPath"`
	Money        int64  `json:"money""`
}

func BuildOrder(order *model.Order, product *model.Product, address *model.Address, host string, port string) Order {
	return Order{
		Id:           int(order.ID),
		OrderNum:     order.OrderNum,
		CreatedAt:    order.CreatedAt.Unix(),
		UpdatedAt:    order.UpdatedAt.Unix(),
		UserId:       order.UserId,
		ProductId:    product.ID,
		BossId:       product.BossId,
		Num:          product.Num,
		AddressName:  address.Name,
		AddressPhone: address.Phone,
		Address:      address.Address,
		Type:         order.Type,
		ProductName:  product.Name,
		ImgPath:      fmt.Sprintf("http://%v:%v/product/Image/%v", host, port, product.ImgPath),
		Money:        int64(order.Money),
	}
}
