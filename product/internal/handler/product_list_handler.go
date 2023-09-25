package handler

import (
	"net/http"

	"E-commerce_system/product/internal/logic"
	"E-commerce_system/product/internal/svc"
	"E-commerce_system/product/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ProductListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductListRequset
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewProductListLogic(r.Context(), svcCtx)
		resp, err := l.ProductList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
