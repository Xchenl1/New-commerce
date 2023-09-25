package logic

import (
	"E-commerce_system/code"
	"E-commerce_system/model"
	"E-commerce_system/product/internal/dao"
	"E-commerce_system/product/internal/svc"
	"E-commerce_system/product/internal/types"
	"E-commerce_system/serializer"
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductSelectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductSelectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductSelectLogic {
	return &ProductSelectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductSelectLogic) ProductSelect(req *types.ProductSelectRequst) (resp *types.ResponseList, err error) {
	var page model.BasePage
	resp = new(types.ResponseList)
	page.PageSize = 10
	page.PageNum = 1
	var product []*model.Product
	total, err := dao.SelectProduct(l.svcCtx.DB, req.Info, &product, page)
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
	resp.Total = int(total)
	resp.Data = productlist
	return
}
