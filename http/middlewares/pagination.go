package middlewares

import (
	"strconv"

	"github.com/stdyum/api-common/databases/pagination"
	"github.com/stdyum/api-common/hc"
)

func PaginationMiddleware(defaultPerPage int) hc.HandlerFunc {
	return func(ctx *hc.Context) {

		cursor := ctx.Query("cursor")
		search := ctx.Query("search")
		pageStr := ctx.Query("page")
		perPageStr := ctx.Query("perPage")
		order := ctx.Query("order")

		page, _ := strconv.Atoi(pageStr)
		if page <= 0 {
			page = 1
		}

		perPage, _ := strconv.Atoi(perPageStr)
		if perPage <= 0 {
			perPage = defaultPerPage
		}

		paginationQuery := pagination.CreatedAtPageQuery{
			QCursor:  cursor,
			QSearch:  search,
			QPage:    page,
			QPerPage: perPage,
			QOrder:   order,
		}

		ctx.Set("pagination", paginationQuery)
		ctx.Next()
	}
}
