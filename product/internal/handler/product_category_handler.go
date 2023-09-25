package handler

import (
	"net/http"

	"E-commerce_system/product/internal/logic"
	"E-commerce_system/product/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ProductCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewProductCategoryLogic(r.Context(), svcCtx)
		resp, err := l.ProductCategory()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
