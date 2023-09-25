package logic

import (
	"E-commerce_system/cart/internal/dao"
	"E-commerce_system/cart/internal/svc"
	"E-commerce_system/cart/internal/types"
	"E-commerce_system/code"
	"E-commerce_system/model"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartUpdateLogic {
	return &CartUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartUpdateLogic) CartUpdate(req *types.CartUpdateRequest) (resp *types.Response, err error) {
	//todo:更新根据id更新购物车
	resp = new(types.Response)
	var cart model.Cart
	err = dao.GetCartById(l.svcCtx.DB, req.Id, &cart)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		return
	}
	if cart.ID <= 0 {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = "未找到购物车"
		return
	}
	cart.Num = uint(req.Num)
	err = dao.UpdateCartNumById(l.svcCtx.DB, req.Id, &cart)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = "更新失败！"
		return
	}
	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	resp.Data = "更新成功！"
	return
}
