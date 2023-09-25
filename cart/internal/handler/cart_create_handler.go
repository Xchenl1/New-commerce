package handler

import (
	"net/http"

	"E-commerce_system/cart/internal/logic"
	"E-commerce_system/cart/internal/svc"
	"E-commerce_system/cart/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CartCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CartCreateRequset
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCartCreateLogic(r.Context(), svcCtx)
		resp, err := l.CartCreate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
