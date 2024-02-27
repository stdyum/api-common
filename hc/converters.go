package hc

import (
	"github.com/gin-gonic/gin"
)

func ConvertHandleFuncFromGin(handlerFunc gin.HandlerFunc) HandlerFunc {
	return func(context *Context) {
		handlerFunc(context.Context)
	}
}

func ConvertHandleFuncToGin(handlerFunc HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		handlerFunc(&Context{context})
	}
}

func ConvertHandleFuncArrayToGin(handlerFunc []HandlerFunc) []gin.HandlerFunc {
	ginHandlers := make([]gin.HandlerFunc, len(handlerFunc))
	for i, handler := range handlerFunc {
		ginHandlers[i] = ConvertHandleFuncToGin(handler)
	}

	return ginHandlers
}

func ConvertIRoutesFromGin(routes gin.IRoutes) IRoutes {
	return &Routes{g: routes}
}

func ConvertGroupFromGin(group *gin.RouterGroup) *RouterGroup {
	return &RouterGroup{
		IRoutes: ConvertIRoutesFromGin(group),
		g:       *group,
	}
}
