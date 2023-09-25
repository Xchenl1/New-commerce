package logic

import (
	"E-commerce_system/code"
	"E-commerce_system/encrypt"
	"E-commerce_system/model"
	"E-commerce_system/order/internal/dao"
	"context"
	"encoding/json"
	"strconv"

	"E-commerce_system/order/internal/svc"
	"E-commerce_system/order/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayOrderLogic {
	return &PayOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayOrderLogic) PayOrder(req *types.PayOrderRequest) (resp *types.Response, err error) {
	// todo: 首先查询商品是否存在
	resp = new(types.Response)
	var product model.Product
	tx := l.svcCtx.DB.Begin()
	//fmt.Println(tx)
	err = dao.GetProductById(tx, req.Product_id, &product)
	if err != nil {
		tx.Rollback()
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = err
		return
	}
	//todo:解密用户金额
	encrypt.Encrypt.SetKey(l.ctx.Value("key").(string)) //获取密钥
	var user model.User
	//fmt.Println(user)
	userid, _ := l.ctx.Value("userId").(json.Number).Float64()
	err = dao.GetUserById(tx, uint(userid), &user)
	if err != nil {
		tx.Rollback()
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = err
		return
	}
	key := l.ctx.Value("key").(string)
	encrypt.Encrypt.SetKey(key)
	money := encrypt.Encrypt.AesDecoding(user.Money)
	money1, _ := strconv.Atoi(money)
	//fmt.Println(money1)
	if money1-req.Money < 0.0 {
		tx.Rollback()
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = "余额不足"
		return
	}
	//todo:重新加密更新用户金额
	user.Money = strconv.Itoa(money1 - req.Money)
	user.Money = encrypt.Encrypt.AesEncoding(user.Money)
	err = dao.UpdateUser(tx, int(user.ID), &user)
	if err != nil {
		tx.Rollback()
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = "更新失败！"
		return
	}
	//todo:更新订单状态 改为0：已支付
	var order model.Order
	err = dao.GetOrderById(tx, req.Order_id, req.OrderNum, &order)
	if err != nil {
		tx.Rollback()
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = "查询订单出错！"
		return
	}
	//fmt.Println(order)
	order.Type = 0
	err = dao.UpdateOrder(tx, int(order.ID), int(order.OrderNum), &order)

	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	resp.Data = "支付成功！"
	tx.Commit()
	return
}
