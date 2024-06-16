package pagination

import (
	"encoding/hex"

	"github.com/stdyum/api-common/entities"
)

func GetNextCreatedAtQuery(rows []entities.Timed, total int, query Query) *CreatedAtPageQuery {
	if query.Page()*query.PerPage() >= total || len(rows) == 0 {
		return nil
	}

	createdAt := rows[len(rows)-1].CreatedAt
	time, _ := createdAt.MarshalBinary()

	return &CreatedAtPageQuery{
		QCursor:  hex.EncodeToString(time),
		QPage:    query.Page() + 1,
		QPerPage: query.PerPage(),
		QOrder:   "aes",
	}
}

func GetPreviousCreatedAtQuery(rows []entities.Timed, query Query) *CreatedAtPageQuery {
	if query.Page() == 1 || len(rows) == 0 {
		return nil
	}

	createdAt := rows[0].CreatedAt
	time, _ := createdAt.MarshalBinary()

	return &CreatedAtPageQuery{
		QCursor:  hex.EncodeToString(time),
		QPage:    query.Page() - 1,
		QPerPage: query.PerPage(),
		QOrder:   "desc",
	}
}
