package middlewaers

import (
	"github.com/gin-gonic/gin"
	"github.com/stdyum/api-common/http"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}

		code, err := http.MapError(ctx.Errors[0])
		ctx.JSON(code, err)
	}
}
