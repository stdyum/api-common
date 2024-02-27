package middlewares

import (
	"github.com/stdyum/api-common/grpc"
	"github.com/stdyum/api-common/hc"
)

func AuthMiddleware() hc.HandlerFunc {
	return func(ctx *hc.Context) {
		token := ctx.GetHeader("Authorization")

		user, err := grpc.Auth(ctx, token)
		if err != nil {
			_ = ctx.Error(err)
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
	}
}
