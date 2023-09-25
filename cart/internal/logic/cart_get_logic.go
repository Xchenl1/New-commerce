package logic

import (
	"E-commerce_system/cart/internal/svc"
	"E-commerce_system/cart/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartGetLogic {
	return &CartGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartGetLogic) CartGet() (resp *types.Response, err error) {
	//todo:拿到用户id
	//userid := l.ctx.Value("userId")
	//userid1, _ := userid.(json.Number).Int64() //获取用户id
	//resp = new(types.Response)

	return
}
