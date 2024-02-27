package pagination

import (
	timed "github.com/stdyum/api-common/entities"
)

func FromArrayAndAmount[T timed.ITimed, D any](rows []T, total int, query Query, toDTO func(el T) D) Result[D] {
	timedRows := make([]timed.Timed, len(rows))
	items := make([]D, len(rows))
	for i, row := range rows {
		timedRows[i] = timed.Timed{
			CreatedAt: row.GetCreatedAt(),
			UpdatedAt: row.GetUpdatedAt(),
		}
		items[i] = toDTO(row)
	}

	return Result[D]{
		Items:    items,
		PerPage:  query.PerPage(),
		Total:    total,
		Page:     query.Page(),
		Next:     GetNextCreatedAtQuery(timedRows, total, query),
		Previous: GetPreviousCreatedAtQuery(timedRows, query),
	}
}
