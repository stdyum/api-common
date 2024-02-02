package http

import "github.com/gin-gonic/gin"

type Routes interface {
	ConfigureRoutes() *gin.Engine
}
