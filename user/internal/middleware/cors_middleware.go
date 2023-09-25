package middleware

import (
	"net/http"
)

type CorsMiddleware struct {
}

func NewCorsMiddleware() *CorsMiddleware {
	return &CorsMiddleware{}
}

func (m *CorsMiddleware) CorsHandle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 在处理登录请求的路由处理程序中添加以下代码
		// 设置跨域允许的域名
		//fmt.Println("允许域名所有域名通过！")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// 设置允许的请求方法
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		// 设置允许的请求头部
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		next(w, r)
	}
}
