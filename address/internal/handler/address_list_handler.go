package handler

import (
	"net/http"

	"E-commerce_system/address/internal/logic"
	"E-commerce_system/address/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddressListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAddressListLogic(r.Context(), svcCtx)
		resp, err := l.AddressList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
