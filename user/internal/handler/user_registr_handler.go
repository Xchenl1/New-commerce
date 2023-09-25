package handler

import (
	"net/http"

	"E-commerce_system/user/internal/logic"
	"E-commerce_system/user/internal/svc"
	"E-commerce_system/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserRegistrHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserRegistrLogic(r.Context(), svcCtx)
		resp, err := l.UserRegistr(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
