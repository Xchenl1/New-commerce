package handler

import (
	"net/http"

	"E-commerce_system/cart/internal/logic"
	"E-commerce_system/cart/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CartGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewCartGetLogic(r.Context(), svcCtx)
		resp, err := l.CartGet()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
