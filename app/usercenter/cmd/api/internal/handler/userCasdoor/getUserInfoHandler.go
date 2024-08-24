package userCasdoor

import (
	"net/http"

	xhttp "github.com/zeromicro/x/http"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/logic/userCasdoor"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/svc"
)

func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := userCasdoor.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo()
		if err != nil {
			// code-data 响应格式
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			// code-data 响应格式
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
