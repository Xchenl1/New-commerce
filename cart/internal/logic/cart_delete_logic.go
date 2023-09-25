package logic

import (
	"E-commerce_system/cart/internal/dao"
	"E-commerce_system/code"
	"context"
	"encoding/json"

	"E-commerce_system/cart/internal/svc"
	"E-commerce_system/cart/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartDeleteLogic {
	return &CartDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartDeleteLogic) CartDelete(req *types.CartDeleteRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	userid := l.ctx.Value("userId")
	userid1, _ := userid.(json.Number).Int64()

	err = dao.DeleteCartById(l.svcCtx.DB, uint(req.Id), uint(userid1))
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
