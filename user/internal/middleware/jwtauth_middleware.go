package middleware

import "net/http"

type JwtAuthMiddleware struct {
	JWTAuth string
}

func NewJwtAuthMiddleware() *JwtAuthMiddleware {
	return &JwtAuthMiddleware{}
}

func (m *JwtAuthMiddleware) JwtAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token1 := r.Header.Get("Authorization")
		m.JWTAuth = token1
		next(w, r)
	}
}
