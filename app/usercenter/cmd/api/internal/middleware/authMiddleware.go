package middleware

import (
	"context"
	"crypto/rsa"
	"fmt"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/config"
	"go_zero_dashboard_base/common/errorx"
	"go_zero_dashboard_base/common/globalkey"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

const jwtExpiryTimeout = 60 * time.Second

type JwtHandler struct {
	rsaPublicKey  *rsa.PublicKey
	casdoorClient *casdoorsdk.Client
	Config        config.Config
}
type (
	UnauthorizedCallback func(w http.ResponseWriter, r *http.Request, err error)
)

const (
	jwtAudience    = "aud"
	jwtExpire      = "exp"
	jwtId          = "jti"
	jwtIssueAt     = "iat"
	jwtIssuer      = "iss"
	jwtNotBefore   = "nbf"
	jwtSubject     = "sub"
	noDetailReason = "no detail reason"
)

func NewJwtHandler(casdoorClient *casdoorsdk.Client, config config.Config) (*JwtHandler, error) {
	rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(casdoorClient.Certificate))
	if err != nil {
		return nil, err
	}
	return &JwtHandler{
		rsaPublicKey:  rsaPublicKey,
		casdoorClient: casdoorClient,
		Config:        config,
	}, nil
}

func (handler *JwtHandler) keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return handler.rsaPublicKey, nil
}

func (handler *JwtHandler) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			strToken string
			claims   jwt.RegisteredClaims
		)
		if auth := r.Header.Get("Authorization"); strings.HasPrefix(auth, "Bearer ") {
			strToken = strings.TrimPrefix(auth, "Bearer ")
		}
		if len(strToken) == 0 {
			httpx.Error(w, errorx.NewDefaultError(errorx.AuthErrorCode))
			return
		}

		// 我们明确设置为仅允许 RS256，并且还禁用
		// 声明检查：RegisteredClaims 内部要求 'iat' 到
		// 不晚于 'now'，但我们允许一点漂移。 todo 检查这里是否已经能验证时间
		token, err := jwt.ParseWithClaims(strToken, &claims, handler.keyFunc,
			jwt.WithValidMethods([]string{"RS256"}))
		if err != nil {
			//未生效过期等验证失败
			ds := errorx.NewDefaultError(errorx.AuthErrorCode)
			xhttp.JsonBaseResponse(w, ds)
			return
		}

		ctx := r.Context()
		// 循环 token.Claims 写入ctx
		logx.Infof("token claims: %+v", token.Claims)
		// 设置用户ID的key
		ctx = context.WithValue(ctx, globalkey.SysJwtUserId, claims.Subject)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func detailAuthLog(r *http.Request, reason string) {
	// discard dump error, only for debug purpose
	details, _ := httputil.DumpRequest(r, true)
	logx.Errorf("authorize failed: %s\n=> %+v", reason, string(details))
}
