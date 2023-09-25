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

type UserLoginLogic struct {
	logx.Logger //日志文件
	ctx         context.Context
	svcCtx      *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginRequest) (resp *types.Response, err error) {
	resp = new(types.Response)
	var user model.User
	//查看用户是否存在
	bol, err := dao.ExistOrNotByUsername(l.svcCtx.DB, req.Username, &user)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		return resp, err
	} else if !bol {
		resp.Status = code.ErrorFailEncryption
		resp.Msg = code.GetMsg(code.ErrorExistUserNotFound) //用户不存在
		resp.Data = "用户不存在，请注册！"
		return resp, err
	}
	//密码是否正确
	if bol = user.CheckPassword(req.Password); !bol {
		resp.Status = code.ErrorNotCompare
		resp.Msg = code.GetMsg(code.ErrorNotCompare)
		resp.Data = "密码错误，请重新登录！"
		return resp, err
	}
	//分发token
	atoken, err := tools.GetToken(l.svcCtx.Config.Auth.AccessSecret, user.Username, int64(user.Model.ID), req.Key)
	if err != nil {
		resp.Status = code.ErrorAuthToken
		resp.Msg = code.GetMsg(code.ErrorAuthToken)
		resp.Data = "token 无效"
		return resp, err
	}
	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	resp.Data = fmt.Sprintf("token:%v", atoken)
	return
}
