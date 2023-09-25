package handler

import (
	"net/http"

	"E-commerce_system/user/internal/logic"
	"E-commerce_system/user/internal/svc"
	"E-commerce_system/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserImageUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserImageUplaodRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewUserImageUploadLogic(r.Context(), svcCtx)
		//取出头部的file
		err := r.ParseMultipartForm(32 << 20) // 设置最大内存大小为32MB
		if err != nil {
			httpx.Error(w, err)
			return
		}
		file, _, _ := r.FormFile("file")
		l.SvcCtx.File = file
		resp, err := l.UserImageUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
