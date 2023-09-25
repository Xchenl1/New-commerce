package handler

import (
	"net/http"

	"E-commerce_system/user/internal/logic"
	"E-commerce_system/user/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserGetidHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserGetidLogic(r.Context(), svcCtx)
		resp, err := l.UserGetid()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
