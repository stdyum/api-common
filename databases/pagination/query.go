package pagination

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
)

func AppendPagination(sql string, query Query, varStart int) (string, []any) {
	builder := strings.Builder{}
	builder.WriteString(sql)

	if strings.Contains(sql, "WHERE") {
		builder.WriteString(" AND (")
	} else {
		builder.WriteString(" WHERE (")
	}

	builder.WriteString(query.Field())
	builder.WriteString(" ")

	if query.Order() == "desc" {
		builder.WriteString("<")
	} else {
		builder.WriteString(">")
	}

	builder.WriteString(" $")
	builder.WriteString(strconv.Itoa(varStart))
	builder.WriteString(") ORDER BY ")
	builder.WriteString(query.Field())
	builder.WriteString(" LIMIT ")
	builder.WriteString(strconv.Itoa(query.PerPage()))

	return builder.String(), []any{query.Cursor()}
}

func QueryPaginationContext(ctx context.Context, database *sql.DB, query string, countQuery string, pagination Query, args ...any) (*sql.Rows, int, error) {
	paginationQuery, vars := AppendPagination(query, pagination, len(args)+1)
	result, err := database.QueryContext(ctx, paginationQuery, append(args, vars...)...)
	if err != nil {
		return nil, 0, err
	}

	var total int
	row := database.QueryRowContext(ctx, countQuery, args...)
	if err = row.Scan(&total); err != nil {
		return nil, 0, err
	}

	return result, total, nil
}

func QueryPagination(database *sql.DB, query string, countQuery string, pagination Query, args ...any) (*sql.Rows, int, error) {
	return QueryPaginationContext(context.Background(), database, query, countQuery, pagination, args...)
}
