package userCasdoor

import (
	"context"
	"go_zero_dashboard_base/common/errorx"
	"go_zero_dashboard_base/common/utils"

	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/svc"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPasswordLogic {
	return &UpdateUserPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// casdoor
func (l *UpdateUserPasswordLogic) UpdateUserPassword(req *types.UpdatePasswordReq) error {
	userId := utils.GetCasdoorUserId(l.ctx)
	user, err := l.svcCtx.CasdoorClient.GetUserByUserId(userId)
	_, err = l.svcCtx.CasdoorClient.SetPassword(user.Owner, user.Name, req.OldPassword, req.NewPassword)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return nil
}
