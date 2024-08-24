package middleware

import "net/http"

type CustomJwtHandleMiddleware struct {
}

func NewCustomJwtHandleMiddleware() *CustomJwtHandleMiddleware {
	return &CustomJwtHandleMiddleware{}
}

func (m *CustomJwtHandleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		// 处理请求
		next(w, r)
	}
}
