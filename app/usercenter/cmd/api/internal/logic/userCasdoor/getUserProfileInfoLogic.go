package userCasdoor

import (
	"context"
	"go_zero_dashboard_base/common/errorx"
	"go_zero_dashboard_base/common/utils"

	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/svc"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

// casdoor
func (l *GetUserProfileInfoLogic) GetUserProfileInfo() (resp *types.UserProfileInfoResp, err error) {
	userId := utils.GetCasdoorUserId(l.ctx)
	//从casdoor 中获取资源
	user, err := l.svcCtx.CasdoorClient.GetUserByUserId(userId)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	return &types.UserProfileInfoResp{
		Id:       userId,
		Nickname: user.Name,
		//Email:    user.Email,
		Avatar: user.Avatar,
	}, nil
}
