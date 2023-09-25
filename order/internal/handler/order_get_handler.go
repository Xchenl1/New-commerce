package handler

import (
	"net/http"

	"E-commerce_system/order/internal/logic"
	"E-commerce_system/order/internal/svc"
	"E-commerce_system/order/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OrderGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderGetIdResquest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOrderGetLogic(r.Context(), svcCtx)
		resp, err := l.OrderGet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
