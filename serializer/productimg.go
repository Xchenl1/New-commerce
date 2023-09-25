package serializer

import (
	"E-commerce_system/model"
	"fmt"
)

type ProductImg struct {
	ProductId uint   `json:"productId"`
	ImgPath   string `json:"imgPath"`
}

func BuildProductImg(item *model.ProductImg, host string, port string) ProductImg {
	return ProductImg{
		ProductId: item.ProductId,
		ImgPath:   fmt.Sprintf("http://%v:%v/product/Image/%v", host, port, item.ImgPath),
	}
}

func BuildProductImgs(item []*model.ProductImg, host string, port string) []ProductImg {
	var productimgs []ProductImg
	for _, v := range item {
		productimg := BuildProductImg(v, host, port)
		productimgs = append(productimgs, productimg)
	}
	return productimgs
}
