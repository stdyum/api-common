package middlewaers

import (
	"github.com/gin-gonic/gin"
	"github.com/stdyum/api-common/proto/impl/auth"
	"google.golang.org/grpc"
)

var authGRpcClient auth.AuthClient

func AuthGRpcDefer() func() {
	conn, err := grpc.Dial("")
	if err != nil {
		panic(err)
	}

	authGRpcClient = auth.NewAuthClient(conn)
	return func() { _ = conn.Close() }
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		token := ctx.GetHeader("Authorization")
		claims, err := authGRpcClient.Auth(ctx, &auth.Token{Token: token})
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.Set("user", claims)
	}
}
