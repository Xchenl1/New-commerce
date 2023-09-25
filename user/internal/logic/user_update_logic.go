package logic

import (
	"E-commerce_system/code"
	"E-commerce_system/model"
	"E-commerce_system/tools"
	"E-commerce_system/user/internal/dao"
	"E-commerce_system/user/internal/svc"
	"E-commerce_system/user/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req *types.UserUpdateNickRequest) (resp *types.Response, err error) {
	//验证Token
	var user model.User
	var c int
	resp = new(types.Response)
	userID, name := tools.VerityToken(l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.JwtAuthMiddleware.JWTAuth)
	err = dao.LookUserInfo(l.svcCtx.DB, userID, name, &user) //查看是否存在
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Error = "查询有误！"
		return resp, nil
	}
	fmt.Println(user)
	if user.NickName == req.Nickname { //检查昵称是否重复
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = "昵称重复"
		return
	} else {
		user.NickName = req.Nickname
	}
	err = dao.Update(l.svcCtx.DB, userID, user.NickName, &user)
	if err != nil {
		c = code.Error
		resp.Status = c
		resp.Msg = code.GetMsg(c)
		resp.Error = "更新失败！"
		return resp, err
	}
	fmt.Println(user)
	c = code.Success
	resp.Status = c
	resp.Msg = code.GetMsg(c)
	resp.Data = "更新成功"
	return
}
