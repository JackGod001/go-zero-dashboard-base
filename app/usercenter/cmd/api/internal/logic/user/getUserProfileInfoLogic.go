package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/svc"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/types"
	"go_zero_dashboard_base/common/errorx"
	"go_zero_dashboard_base/common/utils"
)

type GetUserProfileInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserProfileInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProfileInfoLogic {
	return &GetUserProfileInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 常规登陆获取userid 是 int64
func (l *GetUserProfileInfoLogic) GetUserProfileInfo() (resp *types.UserProfileInfoResp, err error) {
	userId := utils.GetCasdoorUserId(l.ctx)
	// todo casdoor
	user, err := l.svcCtx.CasdoorClient.GetUserByUserId(userId)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	// userId 转化为 string
	return &types.UserProfileInfoResp{
		Id:       userId,
		Nickname: user.Name,
		//Email:    user.Email,
		Avatar: user.Avatar,
	}, nil
}
