package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/svc"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/types"
	"go_zero_dashboard_base/common/errorx"
	"go_zero_dashboard_base/common/utils"
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

// 常规登陆方式
func (l *UpdateUserPasswordLogic) UpdateUserPassword(req *types.UpdatePasswordReq) error {
	userId := utils.GetUserId(l.ctx)
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	if user.Password != utils.MD5(req.OldPassword+l.svcCtx.Config.Salt) {
		return errorx.NewDefaultError(errorx.PasswordErrorCode)
	}

	user.Password = utils.MD5(req.NewPassword + l.svcCtx.Config.Salt)
	_, err = l.svcCtx.UserModel.Update(l.ctx, nil, user)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return nil
}
