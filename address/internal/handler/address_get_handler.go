package handler

import (
	"net/http"

	"E-commerce_system/address/internal/logic"
	"E-commerce_system/address/internal/svc"
	"E-commerce_system/address/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddressGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddressGetIdResquest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddressGetLogic(r.Context(), svcCtx)
		resp, err := l.AddressGet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
