package logic

import (
	"E-commerce_system/code"
	"E-commerce_system/model"
	"E-commerce_system/tools"
	"E-commerce_system/user/internal/dao"
	"E-commerce_system/user/internal/svc"
	"E-commerce_system/user/internal/types"
	"context"
	"gopkg.in/mail.v2"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSendEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSendEmailLogic {
	return &UserSendEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserSendEmailLogic) UserSendEmail(req *types.UserSendEmail) (resp *types.Response, err error) {
	resp = new(types.Response)
	userid, username := tools.VerityToken(l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.JwtAuthMiddleware.JWTAuth)
	//获取邮箱验证的token
	token, err := tools.GetSendToken(l.svcCtx.Config.Auth.AccessSecret, username, userid, req.Email, req.Password, req.OperationType)
	if err != nil {
		c := code.ErrorAuthToken
		resp.Status = c
		resp.Msg = code.GetMsg(c)
		resp.Error = "生成token 失败"
		return resp, err
	}
	//获取notice
	var notice model.Notice
	OperationType, _ := strconv.Atoi(req.OperationType)
	err = dao.GetNoticeById(l.svcCtx.DB, int64(OperationType), &notice)
	if err != nil {
		c := code.Error
		resp.Status = c
		resp.Msg = code.GetMsg(c)
		resp.Error = "notice不存在！"
		return resp, err
	}
	address := l.svcCtx.Config.Email.ValidEmail + token
	mailtext := notice.Text
	mailTex := strings.Replace(mailtext, "email", address, -1) //将所得email替换为address
	m := mail.NewMessage()
	m.SetHeader("From", l.svcCtx.Config.Email.SmtpEmail) //发起方
	m.SetHeader("To", req.Email)                         //接收方
	m.SetHeader("Subject", "FanOne")                     //主题
	m.SetBody("text/html", mailTex)                      //正文内容
	d := mail.NewDialer(l.svcCtx.Config.Email.SmtpHost, 465, l.svcCtx.Config.Email.SmtpEmail, l.svcCtx.Config.Email.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err = d.DialAndSend(m); err != nil {
		c := code.ErrorSendEmail
		resp.Status = c
		resp.Msg = code.GetMsg(c)
		return resp, err
	}
	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	return
}
