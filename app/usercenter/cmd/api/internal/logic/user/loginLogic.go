package user

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/core/logx"
	jwt2 "go_zero_dashboard_base/app/usercenter/cmd/api/internal/common"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/svc"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/types"
	"go_zero_dashboard_base/app/usercenter/model"
	"go_zero_dashboard_base/common/errorx"
	"go_zero_dashboard_base/common/globalkey"
	"go_zero_dashboard_base/common/utils"
	"strconv"
	"time"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}
type GenerateTokenResp struct {
	AccessToken  string
	AccessExpire int64
	RefreshAfter int64
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 常规登陆
func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	verifyCode, err := l.svcCtx.Redis.Get(l.ctx, globalkey.SysLoginCaptchaCachePrefix+req.CaptchaId).Result()
	if err != nil {
		if err == redis.Nil {
			logx.WithContext(l.ctx).Errorf("验证码不存在", err.Error())
			return nil, errorx.NewDefaultError(errorx.CaptchaErrorCode)
		} else {
			logx.WithContext(l.ctx).Errorf("从Redis中获取验证码失败", err.Error())
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
	}
	if verifyCode != req.VerifyCode {
		return nil, errorx.NewDefaultError(errorx.CaptchaErrorCode)
	}
	userId, err := l.loginByEmail(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	//2、生成token
	tokenResp, err := jwt2.NewJwtToken(l.svcCtx, userId)
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

// 根据邮箱和密码登录
func (l *LoginLogic) loginByEmail(email string, password string) (int64, error) {
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, email)
	if err != nil {
		if err != model.ErrNotFound {
			return 0, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
	}
	if user == nil {
		return 0, errorx.NewDefaultError(errorx.UserIdErrorCode)
	}

	if !(utils.Md5ByString(password+l.svcCtx.Config.Salt) == user.Password) {
		return 0, errorx.NewDefaultError(errorx.PasswordErrorCode)
	}

	return user.Id, nil
}
