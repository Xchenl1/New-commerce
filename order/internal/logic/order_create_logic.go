package logic

import (
	"E-commerce_system/code"
	"E-commerce_system/model"
	"E-commerce_system/order/internal/dao"
	"E-commerce_system/order/internal/svc"
	"E-commerce_system/order/internal/types"
	"E-commerce_system/snow"
	"context"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderCreateLogic {
	return &OrderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderCreateLogic) OrderCreate(req *types.OrderCreate) (resp *types.Response, err error) {
	var order model.Order
	resp = new(types.Response)

	//todo:判断地址是否存在
	var address model.Address
	err = dao.GetAddressById(l.svcCtx.DB, req.AddressId, &address)
	if err != nil {
		resp.Status = code.ErrorAddressNotFind
		resp.Msg = code.GetMsg(code.ErrorAddressNotFind)
		return
	}
	req.ProductId = address.ID

	//todo:接下来创建订单编号
	work := snow.NewWorker(5, 5)
	id, err := work.NextID()
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = "生成雪花id错误"
		return
	}
	req.OrderNum = int(id)

	//todo:创建订单
	order = model.Order{
		Model:     gorm.Model{},
		UserId:    req.UserId,
		ProductId: req.ProductId,
		BossId:    req.BossId,
		AddressId: uint(req.AddressId),
		Num:       req.Num,
		OrderNum:  uint64(req.OrderNum),
		Type:      uint(req.Type),
		Money:     float64(req.Money),
	}
	err = dao.CreateOrder(l.svcCtx.DB, &order)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = "创建订单错误"
		return
	}
	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	resp.Data = "创建订单成功！"
	return
}
