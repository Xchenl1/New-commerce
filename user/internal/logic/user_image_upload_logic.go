package logic

import (
	"E-commerce_system/code"
	"E-commerce_system/model"
	"E-commerce_system/service"
	"E-commerce_system/tools"
	"E-commerce_system/user/internal/dao"
	"E-commerce_system/user/internal/svc"
	"E-commerce_system/user/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserImageUploadLogic struct {
	logx.Logger
	ctx    context.Context
	SvcCtx *svc.ServiceContext
}

func NewUserImageUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserImageUploadLogic {
	return &UserImageUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		SvcCtx: svcCtx,
	}
}

func (l *UserImageUploadLogic) UserImageUpload(req *types.UserImageUplaodRequest) (resp *types.Response, err error) {
	resp = new(types.Response)
	userID, name := tools.VerityToken(l.SvcCtx.Config.Auth.AccessSecret, l.SvcCtx.JwtAuthMiddleware.JWTAuth) //验证token
	var user model.User
	err = dao.LookUserInfo(l.SvcCtx.DB, userID, name, &user)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Error = "查询有误！"
		return resp, nil
	}
	//上传文件
	path, err := service.UploadAvatarToLocalStatic(l.SvcCtx.File, userID, name)
	if err != nil {
		c := code.ErrorUploadFail
		resp.Status = c
		resp.Msg = code.GetMsg(c)
		return resp, err
	}
	user.Avatar = path
	err = dao.Update(l.SvcCtx.DB, userID, name, &user)
	if err != nil {
		c := code.Error
		resp.Msg = code.GetMsg(c)
		resp.Status = c
		return resp, err
	}
	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	resp.Data = fmt.Sprintf("http://%v:%v/user/Image/%v", l.SvcCtx.Config.Host, l.SvcCtx.Config.Port, path)
	return
}
