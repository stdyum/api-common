package pagination

import (
	"encoding/hex"
	"time"
)

type CreatedAtPageQuery struct {
	QCursor  string `json:"cursor"`
	QPage    int    `json:"page"`
	QPerPage int    `json:"perPage"`
	QOrder   string `json:"order"`
}

func (c *CreatedAtPageQuery) Field() string {
	return "created_at"
}

func (c *CreatedAtPageQuery) Order() string {
	return c.QOrder
}

func (c *CreatedAtPageQuery) Cursor() any {
	bytes, _ := hex.DecodeString(c.QCursor)
	var t time.Time
	_ = t.UnmarshalBinary(bytes)
	return t
}

func (c *CreatedAtPageQuery) PerPage() int {
	if c.QPerPage == 0 {
		return 10
	}
	return c.QPerPage
}

func (c *CreatedAtPageQuery) Page() int {
	return c.QPage
}

type Pagination struct {
	Items             any                `json:"items"`
	Total             int                `json:"total"`
	Page              int                `json:"page"`
	PerPage           int                `json:"perPage"`
	NextPageQuery     CreatedAtPageQuery `json:"nextPageQuery"`
	PreviousPageQuery CreatedAtPageQuery `json:"previousPageQuery"`
}
