package handler

import (
	"net/http"

	"E-commerce_system/user/internal/logic"
	"E-commerce_system/user/internal/svc"
	"E-commerce_system/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLoginRequest //解析字段
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		//读取自带的配置文件
		l := logic.NewUserLoginLogic(r.Context(), svcCtx)
		//核心是这个函数
		resp, err := l.UserLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
