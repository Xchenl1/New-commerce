package logic

import (
	"E-commerce_system/code"
	"E-commerce_system/model"
	"E-commerce_system/tools"
	"E-commerce_system/user/internal/dao"
	"E-commerce_system/user/internal/svc"
	"E-commerce_system/user/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserGetidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserGetidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserGetidLogic {
	return &UserGetidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserGetidLogic) UserGetid() (resp *types.Response, err error) {
	resp = new(types.Response)
	userID, name := tools.VerityToken(l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.JwtAuthMiddleware.JWTAuth) //验证token
	var user model.User
	err = dao.LookUserInfo(l.svcCtx.DB, userID, name, &user)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Error = "查询有误！"
		return resp, nil
	}
	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	resp.Data = user
	return
}
