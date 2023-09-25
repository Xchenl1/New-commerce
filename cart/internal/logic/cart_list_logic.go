package logic

import (
	"E-commerce_system/cart/internal/dao"
	"E-commerce_system/code"
	"E-commerce_system/model"
	"E-commerce_system/serializer"
	"context"
	"strconv"

	"E-commerce_system/cart/internal/svc"
	"E-commerce_system/cart/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartListLogic {
	return &CartListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartListLogic) CartList(req *types.CartGetIdResquest) (resp *types.Response, err error) {
	// todo: 读取所有的购物车到list中
	var carts []*model.Cart
	resp = new(types.Response)
	err = dao.ListCart(l.svcCtx.DB, &carts)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		return
	}
	var carts1 []*serializer.Cart
	//todo:构造返回response
	for _, v := range carts {
		//todo:检查购物车是否存在
		err = dao.GetCartById(l.svcCtx.DB, int(v.ID), v)
		if err != nil {
			resp.Status = code.ErrorProductNotExist
			resp.Msg = code.GetMsg(code.ErrorExistUserNotFound)
			return
		}
		//todo:检查商品是否存在
		var product model.Product
		_, err = dao.ExistOrNorProduct(l.svcCtx.DB, int(v.ProductId), &product)
		if err != nil {
			resp.Status = code.ErrorProductNotExist
			resp.Msg = code.GetMsg(code.ErrorProductNotExist)
			return
		}
		//todo:检查boss是否存在
		var boss model.User
		err = dao.GetUserInfo(l.svcCtx.DB, int64(v.BossId), &boss)
		if err != nil {
			resp.Status = code.ErrorBossNotExist
			resp.Msg = code.GetMsg(code.ErrorBossNotExist)
			return
		}
		cart := serializer.BuildCart(v, &product, &boss, l.svcCtx.Config.Host, strconv.Itoa(l.svcCtx.Config.Port))
		carts1 = append(carts1, &cart)
	}
	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	resp.Data = carts1
	return
}
