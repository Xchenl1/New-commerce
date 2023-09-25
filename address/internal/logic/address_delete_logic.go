package logic

import (
	"E-commerce_system/address/internal/dao"
	"E-commerce_system/address/internal/svc"
	"E-commerce_system/address/internal/types"
	"E-commerce_system/code"
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddressDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddressDeleteLogic {
	return &AddressDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddressDeleteLogic) AddressDelete(req *types.AddressDeleteRequest) (resp *types.Response, err error) {
	//fmt.Println(req)
	userid := l.ctx.Value("userId")
	userid1, _ := userid.(json.Number).Int64()
	//fmt.Println(userid1)
	err = dao.DeleteAddressById(l.svcCtx.DB, uint(req.Id), uint(userid1))
	resp = new(types.Response)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = err
		return
	}
	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	resp.Data = "成功删除！"
	return
}
