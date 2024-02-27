package middlewares

import (
	"github.com/stdyum/api-common/hc"
	"github.com/stdyum/api-common/http"
)

type errMeta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func ErrorMiddleware() hc.HandlerFunc {
	return func(ctx *hc.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}

		errors := make([]errMeta, len(ctx.Errors))
		code := 0
		for i, ctxError := range ctx.Errors {
			errCode, ctxErr := http.MapError(ctxError)
			errors[i] = errMeta{
				Message: ctxErr.Error(),
				Code:    errCode,
			}

			if code < errCode {
				code = errCode
			}
		}

		ctx.JSON(code, errors)
	}
}
