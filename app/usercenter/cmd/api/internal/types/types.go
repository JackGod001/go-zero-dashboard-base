// Code generated by goctl. DO NOT EDIT.
package types

type GetTokenByCodeReq struct {
	Code  string `json:"code"  label:"code"`
	State string `json:"state"  label:"state"`
}

type GetTokenByCodeResp struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ExpiresAt    int64  `json:"expires_at,omitempty"`
}

type LoginCaptchaResp struct {
	CaptchaId  string `json:"captchaId"`
	VerifyCode string `json:"verifyCode"`
}

type LoginReq struct {
	Email      string `json:"email"     validate:"required,email"         label:"邮箱"`
	Password   string `json:"password"  validate:"required,min=6,max=12"  label:"密码"`
	CaptchaId  string `json:"captchaId"   validate:"required" label:"验证码id"`
	VerifyCode string `json:"verifyCode"  validate:"required" label:"验证码"`
}

type LoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type PageReq struct {
	Page     int64 `json:"page"  validate:"number,gte=1"  label:"页数"`
	PageSize int64 `json:"pageSize" validate:"number,gte=1"  label:"条数"`
}

type Pagination struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
	Total int64 `json:"total"`
}

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token"  validate:"required"  label:"refresh_token"`
}

type RegisterReq struct {
	Email    string `json:"email"     validate:"required,email,max=50"         label:"邮箱"`
	Password string `json:"newPassword"  validate:"required,min=6,max=12"  label:"密码"`
}

type RegisterResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type UpdatePasswordReq struct {
	OldPassword string `json:"oldPassword"  validate:"required,min=6,max=12"  label:"旧密码"`
	NewPassword string `json:"newPassword"  validate:"required,min=6,max=12"  label:"新密码"`
}

type User struct {
	Id       int64  `json:"id"`
	Account  string `json:"account"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Info     string `json:"info"`
}

type UserInfoResp struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type UserProfileInfoResp struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}
