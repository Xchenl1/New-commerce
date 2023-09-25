package handler

import (
	"net/http"

	"E-commerce_system/product/internal/logic"
	"E-commerce_system/product/internal/svc"
	"E-commerce_system/product/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ProductImgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductImgRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewProductImgLogic(r.Context(), svcCtx)
		resp, err := l.ProductImg(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
