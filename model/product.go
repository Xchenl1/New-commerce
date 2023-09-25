package model

import (
	"E-commerce_system/cache"
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"strconv"
)

type Product struct {
	gorm.Model
	Name          string
	CategoryId    uint   //商品类型
	Title         string //商品标题
	Info          string //商品基本信息
	ImgPath       string //图片路径
	Price         string //价格
	Pricediscount string //打折价格
	OnSale        bool   `gorm:"default:false"`
	Num           int    //数量
	BossId        uint   //商家是谁
	BossName      string //商家名字
	BossAvatar    string //图片
}

func (user *Product) TableName() string {
	return "Product"
}

func (Product *Product) View(context context.Context, redis *redis.Client) uint64 {
	countStr, _ := redis.Get(context, cache.ProductViewKey(Product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (Product *Product) AddView(context context.Context, redis *redis.Client) {
	//添加商品的点击数 加1
	redis.Incr(context, cache.ProductViewKey(Product.ID))
	redis.ZIncrBy(context, cache.RanKey, 1, strconv.Itoa(int(Product.ID)))
}
