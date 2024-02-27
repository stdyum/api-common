package hc

import (
	"github.com/gin-gonic/gin"
)

type Routes struct {
	g gin.IRoutes
}

func (r *Routes) Use(handlerFunc ...HandlerFunc) IRoutes {
	ginHandlers := ConvertHandleFuncArrayToGin(handlerFunc)
	ginRoutes := r.g.Use(ginHandlers...)
	return ConvertIRoutesFromGin(ginRoutes)
}

func (r *Routes) Handle(httpMethod, relativePath string, handlerFunc ...HandlerFunc) IRoutes {
	ginHandlers := ConvertHandleFuncArrayToGin(handlerFunc)
	ginRoutes := r.g.Handle(httpMethod, relativePath, ginHandlers...)
	return ConvertIRoutesFromGin(ginRoutes)
}

func (r *Routes) Any(relativePath string, handlerFunc ...HandlerFunc) IRoutes {
	ginHandlers := ConvertHandleFuncArrayToGin(handlerFunc)
	ginRoutes := r.g.Any(relativePath, ginHandlers...)
	return ConvertIRoutesFromGin(ginRoutes)
}

func (r *Routes) GET(relativePath string, handlerFunc ...HandlerFunc) IRoutes {
	ginHandlers := ConvertHandleFuncArrayToGin(handlerFunc)
	ginRoutes := r.g.GET(relativePath, ginHandlers...)
	return ConvertIRoutesFromGin(ginRoutes)
}

func (r *Routes) POST(relativePath string, handlerFunc ...HandlerFunc) IRoutes {
	ginHandlers := ConvertHandleFuncArrayToGin(handlerFunc)
	ginRoutes := r.g.POST(relativePath, ginHandlers...)
	return ConvertIRoutesFromGin(ginRoutes)
}

func (r *Routes) DELETE(relativePath string, handlerFunc ...HandlerFunc) IRoutes {
	ginHandlers := ConvertHandleFuncArrayToGin(handlerFunc)
	ginRoutes := r.g.DELETE(relativePath, ginHandlers...)
	return ConvertIRoutesFromGin(ginRoutes)
}

func (r *Routes) PATCH(relativePath string, handlerFunc ...HandlerFunc) IRoutes {
	ginHandlers := ConvertHandleFuncArrayToGin(handlerFunc)
	ginRoutes := r.g.PATCH(relativePath, ginHandlers...)
	return ConvertIRoutesFromGin(ginRoutes)
}

func (r *Routes) PUT(relativePath string, handlerFunc ...HandlerFunc) IRoutes {
	ginHandlers := ConvertHandleFuncArrayToGin(handlerFunc)
	ginRoutes := r.g.PUT(relativePath, ginHandlers...)
	return ConvertIRoutesFromGin(ginRoutes)
}

func (r *Routes) OPTIONS(relativePath string, handlerFunc ...HandlerFunc) IRoutes {
	ginHandlers := ConvertHandleFuncArrayToGin(handlerFunc)
	ginRoutes := r.g.OPTIONS(relativePath, ginHandlers...)
	return ConvertIRoutesFromGin(ginRoutes)
}

func (r *Routes) HEAD(relativePath string, handlerFunc ...HandlerFunc) IRoutes {
	ginHandlers := ConvertHandleFuncArrayToGin(handlerFunc)
	ginRoutes := r.g.HEAD(relativePath, ginHandlers...)
	return ConvertIRoutesFromGin(ginRoutes)
}
