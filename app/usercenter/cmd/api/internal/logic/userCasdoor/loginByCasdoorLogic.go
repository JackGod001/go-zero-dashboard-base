package userCasdoor

import (
	"context"
	"go_zero_dashboard_base/common/errorx"
	"go_zero_dashboard_base/common/globalkey"
	"time"

	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/svc"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByCasdoorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginByCasdoorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByCasdoorLogic {
	return &LoginByCasdoorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginByCasdoorLogic) LoginByCasdoor(req *types.GetTokenByCodeReq) (resp *types.GetTokenByCodeResp, err error) {
	// todo: add your logic here and delete this line

	token, err := l.svcCtx.CasdoorClient.GetOAuthToken(req.Code, req.State)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据code state 获取token失败", err.Error())
		return nil, errorx.NewDefaultError(errorx.UserIdErrorCode)
	}
	//验证token时间是否过期,是否正确,这里注意本地时间要与casdoor服务器时间一致否则会出现
	claims, err := l.svcCtx.CasdoorClient.ParseJwtToken(token.AccessToken)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("解析token.AccessToken 失败", err.Error())
		return nil, errorx.NewDefaultError(errorx.UserIdErrorCode)
	}
	//设置登陆用户id到redis
	err = l.svcCtx.Redis.SetEx(l.ctx, globalkey.SysOnlineUserCachePrefix+claims.User.Id, "1", 5*time.Minute).Err()
	if err != nil {
		logx.WithContext(l.ctx).Errorf("设置登陆用户id到redis失败", err.Error())
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}
	return &types.GetTokenByCodeResp{
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		//转化成string
		//Scope:     strconv.Itoa(claims.User.Score),
		ExpiresAt: token.Expiry.Unix(),
	}, nil
}
