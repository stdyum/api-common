package hc

import (
	"github.com/gin-gonic/gin"
)

func New() *Engine {
	return &Engine{gin.New()}
}

func Recovery() HandlerFunc {
	return ConvertHandleFuncFromGin(gin.Recovery())
}

func Logger() HandlerFunc {
	return ConvertHandleFuncFromGin(gin.Logger())
}

func (e *Engine) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	ginHandlers := ConvertHandleFuncArrayToGin(handlers)
	group := e.Engine.Group(relativePath, ginHandlers...)
	return ConvertGroupFromGin(group)
}

func (e *Engine) Use(middleware ...HandlerFunc) IRoutes {
	ginHandlers := ConvertHandleFuncArrayToGin(middleware)
	routes := e.Engine.Use(ginHandlers...)
	return ConvertIRoutesFromGin(routes)
}

func (r *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	ginHandlers := ConvertHandleFuncArrayToGin(handlers)
	group := r.g.Group(relativePath, ginHandlers...)
	return ConvertGroupFromGin(group)
}
