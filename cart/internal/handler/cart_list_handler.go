package handler

import (
	"net/http"

	"E-commerce_system/cart/internal/logic"
	"E-commerce_system/cart/internal/svc"
	"E-commerce_system/cart/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CartListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CartGetIdResquest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCartListLogic(r.Context(), svcCtx)
		resp, err := l.CartList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
