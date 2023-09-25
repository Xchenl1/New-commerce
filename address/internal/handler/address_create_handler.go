package handler

import (
	"net/http"

	"E-commerce_system/address/internal/logic"
	"E-commerce_system/address/internal/svc"
	"E-commerce_system/address/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddressCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddressCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddressCreateLogic(r.Context(), svcCtx)
		resp, err := l.AddressCreate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
