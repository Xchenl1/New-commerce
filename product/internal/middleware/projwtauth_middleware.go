package middleware

import "net/http"

type ProjwtAuthMiddleware struct {
	Jwtauth string
}

func NewProjwtAuthMiddleware() *ProjwtAuthMiddleware {
	return &ProjwtAuthMiddleware{}
}

func (m *ProjwtAuthMiddleware) ProJwtAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token1 := r.Header.Get("Authorization")
		m.Jwtauth = token1
		next(w, r)
	}
}
