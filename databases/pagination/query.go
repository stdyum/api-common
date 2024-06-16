package pagination

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
)

func AppendPagination(sql string, query Query, varStart int, searchColumns []string) (string, []any) {
	args := make([]any, 0)

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
	args = append(args, query.Cursor())
	varStart += 1

	search := query.Search()
	if search != "" {
		builder.WriteString(" AND lower(")
		builder.WriteString(strings.Join(searchColumns, ") || ' ' || lower("))
		builder.WriteString(") LIKE '%' || lower($")
		builder.WriteString(strconv.Itoa(varStart))
		builder.WriteString(") || '%'")
		args = append(args, query.Search())
		varStart += 1
	}

	builder.WriteString(") ORDER BY ")
	builder.WriteString(query.Field())
	builder.WriteString(" LIMIT ")
	builder.WriteString(strconv.Itoa(query.PerPage()))

	return builder.String(), args
}

func AppendCount(sql string, query Query, varStart int, searchColumns []string) (string, []any) {
	args := make([]any, 0)

	builder := strings.Builder{}
	builder.WriteString(sql)

	search := query.Search()
	if search != "" {
		if strings.Contains(sql, "WHERE") {
			builder.WriteString(" AND (")
		} else {
			builder.WriteString(" WHERE (")
		}

		builder.WriteString(" lower(")
		builder.WriteString(strings.Join(searchColumns, ") || ' ' || lower("))
		builder.WriteString(") LIKE '%' || lower($")
		builder.WriteString(strconv.Itoa(varStart))
		builder.WriteString(") || '%'")
		args = append(args, query.Search())
		varStart += 1

		builder.WriteString(")")
	}

	return builder.String(), args
}

func AppendPaginationUncountable(sql string, query Query) (string, []any) {
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

	builder.WriteString(" ?")
	builder.WriteString(") ORDER BY ")
	builder.WriteString(query.Field())
	builder.WriteString(" LIMIT ")
	builder.WriteString(strconv.Itoa(query.PerPage()))

	return builder.String(), []any{query.Cursor()}
}

func QueryPaginationContext(ctx context.Context, database *sql.DB, query string, countQuery string, pagination Query, searchColumns []string, args ...any) (*sql.Rows, int, error) {
	query, vars := AppendPagination(query, pagination, len(args)+1, searchColumns)
	result, err := database.QueryContext(ctx, query, append(args, vars...)...)
	if err != nil {
		return nil, 0, err
	}

	var total int
	countQuery, vars = AppendCount(countQuery, pagination, len(args)+1, searchColumns)
	row := database.QueryRowContext(ctx, countQuery, append(args, vars...)...)
	if err = row.Scan(&total); err != nil {
		return nil, 0, err
	}

	return result, total, nil
}

func QueryPagination(database *sql.DB, query string, countQuery string, pagination Query, searchColumns []string, args ...any) (*sql.Rows, int, error) {
	return QueryPaginationContext(context.Background(), database, query, countQuery, pagination, searchColumns, args...)
}
