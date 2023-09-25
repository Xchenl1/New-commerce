package logic

import (
	"E-commerce_system/code"
	"E-commerce_system/order/internal/dao"
	"E-commerce_system/order/internal/svc"
	"E-commerce_system/order/internal/types"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDeleteLogic {
	return &OrderDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderDeleteLogic) OrderDelete(req *types.DeleteOrderIdRequest) (resp *types.Response, err error) {
	//todo:获取订单id并删除
	resp = new(types.Response)
	err = dao.DeleteOrderById(l.svcCtx.DB, uint(req.OrderId), req.Userid)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = fmt.Sprintf("删除订单失败%v", err)
		return
	}
	//todo:返回成功信息
	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	resp.Data = "删除成功"
	return
}
