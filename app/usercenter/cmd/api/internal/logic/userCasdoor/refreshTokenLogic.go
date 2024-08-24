package userCasdoor

import (
	"context"
	"go_zero_dashboard_base/common/errorx"

	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/svc"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenReq) (resp *types.GetTokenByCodeResp, err error) {

	// 根据refreshToken刷新token
	token, err := l.svcCtx.CasdoorClient.RefreshOAuthToken(req.RefreshToken)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("刷新token失败", err.Error())
		return nil, errorx.NewDefaultError(errorx.UserIdErrorCode)
	}
	return &types.GetTokenByCodeResp{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresAt:    token.Expiry.Unix(),
		TokenType:    token.TokenType,
	}, nil
}
