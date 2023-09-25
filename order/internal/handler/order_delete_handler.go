package handler

import (
	"net/http"

	"E-commerce_system/order/internal/logic"
	"E-commerce_system/order/internal/svc"
	"E-commerce_system/order/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OrderDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteOrderIdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOrderDeleteLogic(r.Context(), svcCtx)
		resp, err := l.OrderDelete(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
