package serializer

import (
	"E-commerce_system/model"
	"fmt"
)

type Cart struct {
	Id            uint   `json:"id"`
	UserId        uint   `json:"userId"`
	ProductId     uint   `json:"productId"`
	CreateAt      int64  `json:"createAt"`
	Num           int    `json:"num"`
	MaxNum        int    `json:"maxNum"`
	Name          string `json:"name"`
	ImgPath       string `json:"imgPath"`
	Check         bool   `json:"check"`
	DiscountPrice string `json:"discountPrice"`
	BossId        uint   `json:"bossId"`
	BossName      string `json:"bossName"`
}

func BuildCart(cart *model.Cart, product *model.Product, user *model.User, host string, port string) Cart {
	return Cart{
		Id:            cart.ID,
		UserId:        cart.UserId,
		ProductId:     cart.ProductId,
		CreateAt:      cart.CreatedAt.Unix(),
		Num:           int(cart.Num),
		MaxNum:        int(cart.MaxNum),
		ImgPath:       fmt.Sprintf("http://%v:%v/product/Image/%v", host, port, product.ImgPath),
		Name:          product.Name,
		Check:         cart.Check,
		DiscountPrice: product.Pricediscount,
		BossId:        user.ID,
		BossName:      user.NickName,
	}
}
