package handler

import (
	"net/http"

	"E-commerce_system/product/internal/logic"
	"E-commerce_system/product/internal/svc"
	"E-commerce_system/product/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		//取出头部的file
		err := r.ParseMultipartForm(32 << 20) // 设置最大内存大小为32MB
		if err != nil {
			httpx.Error(w, err)
			return
		} //添加文件
		form := r.MultipartForm
		files := form.File["image"]
		svcCtx.File = files
		l := logic.NewProductLogic(r.Context(), svcCtx)
		resp, err := l.Product(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
