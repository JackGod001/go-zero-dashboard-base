package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/logic/user"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/svc"
	"go_zero_dashboard_base/common/globalkey"
	"time"
)

// 生成jwttoken
func NewJwtToken(svcCtx *svc.ServiceContext, userId int64) (*user.GenerateTokenResp, error) {
	iat := time.Now().Unix()
	exp := iat + svcCtx.Config.JwtAuth.AccessExpire

	claims := jwt.MapClaims{
		"exp":                  exp,
		"iat":                  iat,
		globalkey.SysJwtUserId: userId,
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	tokenStr, err := token.SignedString([]byte(svcCtx.Config.JwtAuth.AccessSecret))
	if err != nil {
		return nil, err
	}

	return &user.GenerateTokenResp{
		AccessToken:  tokenStr,
		AccessExpire: exp,
		RefreshAfter: iat,
	}, nil
}
