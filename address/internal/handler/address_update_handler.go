package handler

import (
	"net/http"

	"E-commerce_system/address/internal/logic"
	"E-commerce_system/address/internal/svc"
	"E-commerce_system/address/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddressUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddressUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddressUpdateLogic(r.Context(), svcCtx)
		resp, err := l.AddressUpdate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
