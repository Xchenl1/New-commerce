package logic

import (
	"E-commerce_system/address/internal/dao"
	"E-commerce_system/code"
	"E-commerce_system/model"
	"context"
	"encoding/json"

	"E-commerce_system/address/internal/svc"
	"E-commerce_system/address/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddressUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddressUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddressUpdateLogic {
	return &AddressUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddressUpdateLogic) AddressUpdate(req *types.AddressUpdateRequest) (resp *types.Response, err error) {
	userid := l.ctx.Value("userId")
	userid1, _ := userid.(json.Number).Int64() //获取用户id

	var address = &model.Address{
		Address: req.Address,
		Name:    req.Name,
		Phone:   req.Phone,
		UserId:  uint(userid1),
	}
	err = dao.UpdateAddressById(l.svcCtx.DB, req.Id, address)

	resp = new(types.Response)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = err
		return
	}
	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	resp.Data = "修改成功"
	return
}
