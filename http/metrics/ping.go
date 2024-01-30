package metrics

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
}
