package logic

import (
	"E-commerce_system/code"
	"E-commerce_system/encrypt"
	"E-commerce_system/model"
	"E-commerce_system/user/internal/dao"
	"E-commerce_system/user/internal/svc"
	"E-commerce_system/user/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UserRegistrLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegistrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegistrLogic {
	return &UserRegistrLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegistrLogic) UserRegistr(req *types.UserRegisterRequest) (resp *types.Response, err error) {
	resp = new(types.Response)
	var user model.User
	if req.Key == "" {
		resp.Msg = code.GetMsg(code.Error)
		resp.Error = "密钥长度不足"
		resp.Status = code.Error
		return resp, err
	}
	//10000 密文存储  加密操作
	req.Key = string(encrypt.PadKey([]byte(req.Key))) //密钥字节不够 需要添加到16个字节
	fmt.Println(req.Key)
	encrypt.Encrypt.SetKey(req.Key)
	//判断用户是否存在
	bol, err := dao.ExistOrNotByUsername(l.svcCtx.DB, req.Username, &user)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		return nil, err
	}
	if bol {
		resp.Status = code.ErrorExistUser
		resp.Msg = code.GetMsg(code.ErrorExistUser)
		return
	}
	user = model.User{
		Model:    gorm.Model{},
		Username: req.Username,
		Password: "",
		NickName: req.Nickname,
		Status:   model.Active,
		Money:    encrypt.Encrypt.AesEncoding("100000"), //进行加密
	}
	if err = user.SetPassword(req.Password); err != nil {
		resp.Status = code.ErrorFailEncryption
		resp.Msg = code.GetMsg(code.ErrorFailEncryption)
		return
	}
	if err = dao.CreateUser(l.svcCtx.DB, &user); err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		return
	}
	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	return
}
