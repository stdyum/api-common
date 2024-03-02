package metrics

import (
	"github.com/stdyum/api-common/hc"

	"net/http"
)

func ping(ctx *hc.Context) {
	ctx.JSON(http.StatusOK, "pong")
}
