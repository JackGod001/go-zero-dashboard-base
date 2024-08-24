package user

import (
	"context"
	"go_zero_dashboard_base/app/usercenter/model"
	"go_zero_dashboard_base/common/errorx"
	"go_zero_dashboard_base/common/globalkey"
	"go_zero_dashboard_base/common/utils"
	"strconv"
	"time"

	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/svc"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *RegisterLogic) Register(req *types.RegisterReq) error {
	//验证邮箱是否存在
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	if err != nil {
		if err != model.ErrNotFound {
			return errorx.NewDefaultError(errorx.ServerErrorCode)
		}
	}
	if user != nil {
		return errorx.NewDefaultError(errorx.AddUserErrorCode)
	}

	password := utils.Md5ByString(req.Password + l.svcCtx.Config.Salt)
	user = &model.User{
		Email:    req.Email,
		Password: password,
	}
	// 保存数据
	insertResult, err := l.svcCtx.UserModel.Insert(l.ctx, nil, user)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("新增用户插入数据库", err.Error())
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}
	lastId, err = insertResult.LastInsertId()
	if err != nil {
		logx.WithContext(l.ctx).Errorf("新增用户插入数据库获取最后id失败", err.Error())
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}
	//2、生成token
	tokenResp, err := l.getJwtToken(userId)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	// 设置登陆用户id到redis
	err = l.svcCtx.Redis.SetEx(l.ctx, globalkey.SysOnlineUserCachePrefix+strconv.FormatInt(userId, 10), "1", 5*time.Minute).Err()
	return &types.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
