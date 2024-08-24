package user

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/svc"
	"go_zero_dashboard_base/common/globalkey"
	"go_zero_dashboard_base/common/utils"
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

// 常规登陆方式的退出
func (l *LogoutLogic) Logout() error {
	userId := utils.GetUserId(l.ctx)
	//转化为字符串
	userIdStr := fmt.Sprintf("%d", userId)
	err := l.svcCtx.Redis.Del(l.ctx, globalkey.SysOnlineUserCachePrefix+userIdStr).Err()
	if err != nil {
		logx.WithContext(l.ctx).Errorf("删除用户在线状态失败, err: %v", err)
		return err
	}
	return nil
}
