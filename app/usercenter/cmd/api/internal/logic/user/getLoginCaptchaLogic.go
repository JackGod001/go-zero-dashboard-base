package user

import (
	"context"
	"fmt"
	"github.com/mojocn/base64Captcha"
	utils2 "github.com/zeromicro/go-zero/core/utils"
	"go_zero_dashboard_base/common/errorx"
	"go_zero_dashboard_base/common/globalkey"
	"go_zero_dashboard_base/common/utils"
	"time"

	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/svc"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLoginCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLoginCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLoginCaptchaLogic {
	return &GetLoginCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 常规登陆
func (l *GetLoginCaptchaLogic) GetLoginCaptcha() (resp *types.LoginCaptchaResp, err error) {
	// todo: add your logic here and delete this line
	var store = base64Captcha.DefaultMemStore
	captcha := utils.NewCaptcha(45, 80, 4, 40, 30, 89, 0)
	driver := captcha.DriverString()
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := c.Generate()
	if err != nil {
		logx.WithContext(l.ctx).Errorf("生成验证码失败 error:%v ", err.Error())
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}
	val := store.Get(id, true)
	captchaId := utils2.NewUuid()
	err = l.svcCtx.Redis.SetEx(l.ctx, globalkey.SysLoginCaptchaCachePrefix+captchaId, val, 1*time.Minute).Err()
	if err != nil {
		logx.WithContext(l.ctx).Errorf("验证码缓存失败 error:%v ", err.Error())
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}
	fmt.Println("val", val)
	fmt.Println("captchaId", captchaId)
	return &types.LoginCaptchaResp{
		CaptchaId:  captchaId,
		VerifyCode: b64s,
	}, nil
}
