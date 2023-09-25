package logic

import (
	"E-commerce_system/code"
	"E-commerce_system/model"
	"E-commerce_system/order/internal/dao"
	"E-commerce_system/serializer"
	"context"
	"strconv"

	"E-commerce_system/order/internal/svc"
	"E-commerce_system/order/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderGetLogic {
	return &OrderGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderGetLogic) OrderGet(req *types.OrderGetIdResquest) (resp *types.Response, err error) {

	//todo:判断用户是否存在
	resp = new(types.Response)
	count, err := dao.GetOrderUserExistOrNot(l.svcCtx.DB, int(req.UserId))
	if count == 0 {
		resp.Status = code.ErrorExistUserNotFound
		resp.Msg = code.GetMsg(code.ErrorExistUserNotFound)
		return
	}
	//todo:根据订单id和用户id来确定order
	var order model.Order
	err = dao.GetOrderByIdAndUserId(l.svcCtx.DB, req.Id, int(req.UserId), &order)

	//todo:查product
	var product model.Product
	err = dao.GetProductById(l.svcCtx.DB, int(order.ProductId), &product)
	if count == 0 {
		resp.Status = code.ErrorProductNotExist
		resp.Msg = code.GetMsg(code.ErrorProductNotExist)
		return
	}

	//todo:获取address
	var address model.Address
	err = dao.GetAddressById(l.svcCtx.DB, int(order.AddressId), &address)
	if count == 0 {
		resp.Status = code.ErrorAddressNotFind
		resp.Msg = code.GetMsg(code.ErrorAddressNotFind)
		return
	}

	//todo:成功返回
	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	resp.Data = serializer.BuildOrder(&order, &product, &address, l.svcCtx.Config.Host, strconv.Itoa(l.svcCtx.Config.Port))
	return
}
