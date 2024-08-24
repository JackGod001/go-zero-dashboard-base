package userCasdoor

import (
	"context"
	"go_zero_dashboard_base/common/globalkey"
	"go_zero_dashboard_base/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/svc"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// casdoor登出
func (l *LogoutLogic) Logout() error {
	user, err := utils.GetCasdoorUser(l.ctx, l.svcCtx.CasdoorClient)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("获取用户信息失败, err: %v", err)
		return err
	}

	err = l.svcCtx.Redis.Del(l.ctx, globalkey.SysOnlineUserCachePrefix+user.Id).Err()
	if err != nil {
		logx.WithContext(l.ctx).Errorf("删除用户在线状态失败, err: %v", err)
		return err
	}
	return nil
}
