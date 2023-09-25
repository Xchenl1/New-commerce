package logic

import (
	"E-commerce_system/cart/internal/dao"
	"E-commerce_system/code"
	"E-commerce_system/model"
	"E-commerce_system/serializer"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"E-commerce_system/cart/internal/svc"
	"E-commerce_system/cart/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartCreateLogic {
	return &CartCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartCreateLogic) CartCreate(req *types.CartCreateRequset) (resp *types.Response, err error) {
	// todo: 首先判断商品是否存在
	userid := l.ctx.Value("userId")
	userid1, _ := userid.(json.Number).Int64() //获取用户id
	resp = new(types.Response)

	var product model.Product
	bol, err := dao.ExistOrNorProduct(l.svcCtx.DB, req.ProductId, &product) //true：存在用户  false:不存在用户
	//fmt.Println(bol, err)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		return resp, err
	}
	if !bol {
		resp.Status = code.ErrorproductExist
		resp.Msg = code.GetMsg(code.ErrorproductExist)
		return
	}
	fmt.Println(product)

	//todo:创建购物车
	cart := &model.Cart{
		UserId:    uint(userid1),
		ProductId: uint(req.ProductId),
		BossId:    product.BossId, //物品的商家 商家可以用户
		Num:       uint(req.Num),
		MaxNum:    uint(10),
	}
	err = dao.CreateCart(l.svcCtx.DB, cart)
	fmt.Println(cart)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = "创建购物车出错！"
		return
	}

	//todo:获取商家信息
	var user model.User //todo:不能定义*model.user  否则会空指针
	err = dao.GetUserInfo(l.svcCtx.DB, int64(req.BossId), &user)
	fmt.Println(err)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = "查询商家出错！"
		return
	}
	if user.ID <= 0 {
		fmt.Println(user.ID)
		resp.Status = code.ErrorExistUserNotFound
		resp.Msg = code.GetMsg(code.ErrorExistUserNotFound)
		resp.Data = "未发现商家"
		return
	}
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = "查询商家出错！"
		return
	}
	//todo:返回给用户信息
	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	resp.Data = serializer.BuildCart(cart, &product, &user, l.svcCtx.Config.Host, strconv.Itoa(l.svcCtx.Config.Port))
	return
}
