package handler

import (
	"net/http"

	"E-commerce_system/user/internal/logic"
	"E-commerce_system/user/internal/svc"
	"E-commerce_system/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserSendEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserSendEmail
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserSendEmailLogic(r.Context(), svcCtx)
		resp, err := l.UserSendEmail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
