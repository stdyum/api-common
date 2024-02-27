package query_builder

import (
	"strconv"
	"strings"
)

type QueryBuilder struct {
	query string
	args  []any
}

func NewQueryBuilder(query string, args ...any) *QueryBuilder {
	return &QueryBuilder{query: query, args: args}
}

func (q *QueryBuilder) Build() (string, []any) {
	return q.query, q.args
}

func (q *QueryBuilder) Append(query string, args ...any) *QueryBuilder {
	startVarNumber := len(q.args) + 1

	for i := range args {
		indexed := "$" + strconv.Itoa(i+1)
		current := "$" + strconv.Itoa(i+startVarNumber)

		index := strings.Index(query, indexed)
		for index != -1 {
			if query[index+len(indexed)] < 48 || query[index+len(indexed)] > 57 {
				query = query[:index] + current + query[index+len(indexed):]
			}

			previous := index
			index = strings.Index(query[index+1:], indexed)
			if index != -1 {
				index += previous + 1
			}
		}
	}

	q.query += " "
	q.query += query
	q.args = append(q.args, args...)
	return q
}

func (q *QueryBuilder) RemoveLast(i int) {
	q.query = q.query[:len(q.query)-i]
}
