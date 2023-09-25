package logic

import (
	"E-commerce_system/code"
	"E-commerce_system/model"
	"E-commerce_system/tools"
	"E-commerce_system/user/internal/dao"
	"context"
	"strconv"

	"E-commerce_system/user/internal/svc"
	"E-commerce_system/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserValidEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserValidEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserValidEmailLogic {
	return &UserValidEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserValidEmailLogic) UserValidEmail(req *types.UserValidEmail) (resp *types.Response, err error) {
	resp = new(types.Response)
	claim, err := tools.VeritySendToken(l.svcCtx.Config.Auth.AccessSecret, req.Token)
	if err != nil {
		resp.Status = code.ErrorAuthToken
		resp.Msg = code.GetMsg(code.ErrorAuthToken)
		return resp, err
	}
	userid := claim["userId"].(float64)
	username := claim["username"].(string)
	passwd := claim["password"].(string)
	option := claim["OperationType"].(string)
	option1, _ := strconv.Atoi(option)
	email := claim["email"].(string)
	var user model.User
	err = dao.LookUserInfo(l.svcCtx.DB, int64(userid), username, &user) //根据
	if err != nil {
		resp.Status = code.ErrorExistUserNotFound
		resp.Msg = code.GetMsg(code.ErrorExistUser)
		return resp, err
	}
	if option1 == 1 { //注册邮箱
		user.Email = email
	} else if option1 == 2 { //注销邮箱
		user.Email = ""
	} else if option1 == 3 { //修改密码
		err = user.SetPassword(passwd)
		if err != nil {
			resp.Status = code.ErrorFailEncryption
			resp.Msg = code.GetMsg(code.ErrorFailEncryption)
			return resp, err
		}
	}
	//fmt.Println(user)
	err = dao.UpdateEmail(l.svcCtx.DB, int64(userid), user.Email)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Error = "更新失败！"
		return resp, err
	}
	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	return
}
