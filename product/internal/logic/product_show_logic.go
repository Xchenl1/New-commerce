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

type ProductShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductShowLogic {
	return &ProductShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductShowLogic) ProductShow(req *types.ProductShowRequst) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	var product model.Product
	err = dao.ShowProduct(l.svcCtx.DB, req.Id, &product)
	resp = new(types.Response)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = err
		return
	}
	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	resp.Data = serializer.BuildProduct(product, l.svcCtx.Config.Host, strconv.Itoa(l.svcCtx.Config.Port), l.ctx, l.svcCtx.Redis)
	return
}
