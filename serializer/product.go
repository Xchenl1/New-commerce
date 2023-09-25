package serializer

import (
	"E-commerce_system/model"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Product struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	CategoryId    uint   `json:"categoryId"`                  //商品类型
	Title         string `json:"title"`                       //商品标题
	Info          string `json:"info"`                        //商品基本信息
	ImgPath       string `json:"imgPath"`                     //图片路径
	View          int    `json:"view"`                        //点击量
	Price         string `json:"price"`                       //价格
	Pricediscount string `json:"pricediscount"`               //打折价格
	OnSale        bool   `json:"onSale" gorm:"default:false"` //是否在售
	Num           int    `json:"num"`                         //数量
	BossId        uint   `json:"bossId"`                      //商家是谁
	BossName      string `json:"bossName"`                    //商家名字
	BossAvatar    string `json:"bossAvatar"`                  //图片
}

func BuildProduct(item model.Product, host string, port string, context context.Context, redis *redis.Client) Product {
	return Product{
		Id:            item.ID,
		Name:          item.Name,
		CategoryId:    item.CategoryId,
		Title:         item.Title,
		Info:          item.Info,
		ImgPath:       fmt.Sprintf("http://%v:%v/product/Image/%v", host, port, item.ImgPath),
		View:          int(item.View(context, redis)), //获取点击量
		Price:         item.Price,
		Pricediscount: item.Pricediscount,
		OnSale:        item.OnSale,
		Num:           item.Num,
		BossId:        item.BossId,
		BossName:      item.BossName,
		BossAvatar:    fmt.Sprintf("http://%v:%v/product/Image/%v", host, port, item.BossAvatar),
	}
}

func BuildProductList(item *model.Product, host string, port string, context context.Context, redis *redis.Client) *Product {
	return &Product{
		Id:            item.ID,
		Name:          item.Name,
		CategoryId:    item.CategoryId,
		Title:         item.Title,
		Info:          item.Info,
		ImgPath:       fmt.Sprintf("http://%v:%v/product/Image/%v", host, port, item.ImgPath),
		View:          int(item.View(context, redis)), //获取点击量
		Price:         item.Price,
		Pricediscount: item.Pricediscount,
		OnSale:        item.OnSale,
		Num:           item.Num,
		BossId:        item.BossId,
		BossName:      item.BossName,
		BossAvatar:    fmt.Sprintf("http://%v:%v/product/Image/%v", host, port, item.BossAvatar),
	}
}
