package middlewares

import (
	"github.com/stdyum/api-common/grpc"
	"github.com/stdyum/api-common/hc"
)

func EnrollmentAuthMiddleware() hc.HandlerFunc {
	return func(ctx *hc.Context) {
		token := ctx.GetHeader("Authorization")
		studyPlaceId := ctx.Query("studyPlaceId")

		enrollmentUser, err := grpc.EnrollmentAuth(ctx, token, studyPlaceId)
		if err != nil {
			_ = ctx.Error(err)
			ctx.Abort()
			return
		}

		ctx.Set("enrollmentUser", enrollmentUser)
		ctx.Set("enrollment", enrollmentUser.Enrollment)
	}
}
