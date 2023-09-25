package logic

import (
	"E-commerce_system/code"
	"E-commerce_system/model"
	"E-commerce_system/product/internal/dao"
	"E-commerce_system/serializer"
	"context"

	"E-commerce_system/product/internal/svc"
	"E-commerce_system/product/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryLogic {
	return &ProductCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCategoryLogic) ProductCategory() (resp *types.Response, err error) {
	var category []*model.Category
	resp = new(types.Response)
	err = dao.ListProductCategory(l.svcCtx.DB, &category) //获取分类
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = err
		return
	}
	categorys := serializer.BuildCategorys(category)
	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	resp.Data = categorys
	return
}
