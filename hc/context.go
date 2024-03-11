package hc

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stdyum/api-common/databases/pagination"
	"github.com/stdyum/api-common/models"
)

type Context struct {
	*gin.Context
}

func (ctx *Context) UUIDParam(key string) (uuid.UUID, error) {
	id := ctx.Param(key)
	return uuid.Parse(id)
}

func (ctx *Context) QueryTime(key string) (time.Time, error) {
	value := ctx.Query(key)
	return time.Parse(time.RFC3339, value)
}

func (ctx *Context) QueryUUID(key string) (uuid.UUID, error) {
	value := ctx.Query(key)
	return uuid.Parse(value)
}

func (ctx *Context) QueryInt(key string) (int, error) {
	value := ctx.Query(key)
	return strconv.Atoi(value)
}

func (ctx *Context) PaginationQuery() pagination.CreatedAtPageQuery {
	pAny, ok := ctx.Get("pagination")
	if !ok {
		return pagination.CreatedAtPageQuery{}
	}

	if p, ok := pAny.(pagination.CreatedAtPageQuery); ok {
		return p
	}

	return pagination.CreatedAtPageQuery{}
}

func (ctx *Context) User() models.User {
	uAny, ok := ctx.Get("user")
	if !ok {
		return models.User{}
	}

	if u, ok := uAny.(models.User); ok {
		return u
	}

	return models.User{}
}

func (ctx *Context) Enrollment() models.Enrollment {
	uAny, ok := ctx.Get("enrollment")
	if !ok {
		return models.Enrollment{}
	}

	if u, ok := uAny.(models.Enrollment); ok {
		return u
	}

	return models.Enrollment{}
}
