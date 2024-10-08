package middleware

import (
	"context"
	"crypto/rsa"
	"fmt"
	"go_zero_dashboard_base/common/globalkey"
	"net/http"
	"strings"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero_dashboard_base/common/errorx"
)

type CasdoorJwtMiddleware struct {
	rsaPublicKey *rsa.PublicKey
	casdoorsdk   *casdoorsdk.Client
}

func NewCasdoorJwtMiddleware(casdoorClient *casdoorsdk.Client) (*CasdoorJwtMiddleware, error) {
	cerBy := []byte(casdoorClient.Certificate)
	rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(cerBy)
	if err != nil {
		return nil, err
	}
	return &CasdoorJwtMiddleware{
		rsaPublicKey: rsaPublicKey,
		casdoorsdk:   casdoorClient,
	}, nil
}
func (m *CasdoorJwtMiddleware) keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return m.rsaPublicKey, nil
}
func (m *CasdoorJwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			httpx.Error(w, errorx.NewDefaultError(errorx.AuthErrorCode))
			return
		}

		token := strings.Split(authHeader, "Bearer ")
		if len(token) != 2 {
			httpx.Error(w, errorx.NewDefaultError(errorx.AuthErrorCode))
			return
		}
		// 有时与服务器时间不同步会导致token解析失败，优先同步时间
		//time.Sleep(500 * time.Millisecond)
		claims, err := m.casdoorsdk.ParseJwtToken(token[1])
		if err != nil {
			httpx.Error(w, errorx.NewDefaultError(errorx.AuthErrorCode))
			return
		}
		ctx := r.Context()
		//// 设置用户ID的key
		ctx = context.WithValue(ctx, globalkey.SysJwtUserId, claims.Subject)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
