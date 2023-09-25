package logic

import (
	"E-commerce_system/address/internal/dao"
	"E-commerce_system/code"
	"E-commerce_system/model"
	"E-commerce_system/serializer"
	"context"
	"encoding/json"
	"fmt"

	"E-commerce_system/address/internal/svc"
	"E-commerce_system/address/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddressListLogic {
	return &AddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddressListLogic) AddressList() (resp *types.ResponseList, err error) {
	// todo: add your logic here and delete this line
	userid := l.ctx.Value("userId")
	userid1, _ := userid.(json.Number).Int64()
	fmt.Println(userid1)

	addressmap := make(map[string]interface{}, 0)
	addressmap["user_id"] = userid1
	fmt.Println()
	var address []*model.Address
	resp = new(types.ResponseList)
	err, count := dao.GetAddressByUserid(l.svcCtx.DB, addressmap, &address)
	fmt.Println(address)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = err
		return
	}
	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	resp.Data = serializer.BuildAddresses(address)
	resp.Total = int(count)
	return
}
