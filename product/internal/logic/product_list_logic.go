package logic

import (
	"E-commerce_system/code"
	"E-commerce_system/model"
	"E-commerce_system/product/internal/dao"
	"E-commerce_system/serializer"
	"context"
	"fmt"
	"strconv"

	"E-commerce_system/product/internal/svc"
	"E-commerce_system/product/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductListLogic) ProductList(req *types.ProductListRequset) (resp *types.ResponseList, err error) {
	resp = new(types.ResponseList)
	fmt.Println(req)
	var total int64
	condation := make(map[string]interface{})
	if req.CategoryId == 0 {
		req.CategoryId = 1
	}
	if req.PageSize == 0 { //页大小
		req.PageSize = 10
	}
	if req.PageNum == 0 { //第几页
		req.PageNum = 1
	}
	condation["category_id"] = req.CategoryId
	var product []*model.Product
	err = dao.ListProduct(l.svcCtx.DB, condation, req.PageSize, req.PageNum, &product) //计算列表
	//fmt.Println("product:", product)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = err
		return
	}
	err = dao.CountProduct(l.svcCtx.DB, condation, &total) //总数
	//fmt.Println("总数：", total)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = err
		return
	}
	productlist := make([]*serializer.Product, 0)
	for _, v := range product {
		proresp := serializer.BuildProductList(v, l.svcCtx.Config.Host, strconv.Itoa(l.svcCtx.Config.Port), l.ctx, l.svcCtx.Redis)
		//fmt.Println(proresp)
		productlist = append(productlist, proresp)
	}
	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	resp.Data = productlist
	resp.Total = int(total)
	return
}
