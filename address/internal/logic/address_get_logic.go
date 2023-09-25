package logic

import (
	"E-commerce_system/address/internal/dao"
	"E-commerce_system/code"
	"E-commerce_system/model"
	"E-commerce_system/serializer"
	"context"
	"fmt"

	"E-commerce_system/address/internal/svc"
	"E-commerce_system/address/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddressGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddressGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddressGetLogic {
	return &AddressGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddressGetLogic) AddressGet(req *types.AddressGetIdResquest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	var address []*model.Address
	resp = new(types.Response)
	err = dao.GetAddress(l.svcCtx.DB, req.Id, &address)
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
	return
}
