package middlewares

import (
	"errors"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/hc"
)

func StudyPlaceMiddleware() hc.HandlerFunc {
	return func(ctx *hc.Context) {
		studyPlaceIdString := ctx.GetHeader("Study-Place-Id")
		studyPlaceId, err := uuid.Parse(studyPlaceIdString)
		if err != nil {
			_ = ctx.Error(errors.New("invalid study-place-id"))
			ctx.Abort()
			return
		}

		ctx.Set("studyPlaceId", studyPlaceId)
	}
}
