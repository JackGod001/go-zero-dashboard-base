package user

import (
	"errors"
	translations "github.com/go-playground/validator/v10/translations/zh"
	"go_zero_dashboard_base/common/errorx"
	"net/http"
	"reflect"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/logic/user"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/svc"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/types"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	xhttp "github.com/zeromicro/x/http"
)

func UpdateUserPasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdatePasswordReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewHandlerError(errorx.ParamErrorCode, err.Error()))
			return
		}

		validate := validator.New()
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			return name
		})

		trans, _ := ut.New(zh.New()).GetTranslator("zh")
		validateErr := translations.RegisterDefaultTranslations(validate, trans)
		if validateErr = validate.StructCtx(r.Context(), req); validateErr != nil {
			for _, err := range validateErr.(validator.ValidationErrors) {
				httpx.Error(w, errorx.NewHandlerError(errorx.ParamErrorCode, errors.New(err.Translate(trans)).Error()))
				return
			}
		}

		l := user.NewUpdateUserPasswordLogic(r.Context(), svcCtx)
		err := l.UpdateUserPassword(&req)
		if err != nil {
			// code-data 响应格式
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			// code-data 响应格式
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
