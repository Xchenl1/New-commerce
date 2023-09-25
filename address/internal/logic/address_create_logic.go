package logic

import (
	"E-commerce_system/address/internal/dao"
	"E-commerce_system/code"
	"E-commerce_system/model"
	"E-commerce_system/serializer"
	"context"
	"encoding/json"
	"gorm.io/gorm"

	"E-commerce_system/address/internal/svc"
	"E-commerce_system/address/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddressCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddressCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddressCreateLogic {
	return &AddressCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddressCreateLogic) AddressCreate(req *types.AddressCreateRequest) (resp *types.Response, err error) {
	userid := l.ctx.Value("userId")
	userid1, _ := userid.(json.Number).Int64()

	var address = model.Address{
		Model:   gorm.Model{},
		UserId:  uint(userid1),
		Name:    req.Name,
		Phone:   req.Phone,
		Address: req.Address,
	}
	err = dao.CreateAddress(l.svcCtx.DB, &address)
	resp = new(types.Response)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = err
		return
	}
	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	resp.Data = serializer.BuildAddress(&address)
	return
}
