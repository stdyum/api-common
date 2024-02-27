package pagination

type Result[T any] struct {
	Items    []T                 `json:"items"`
	PerPage  int                 `json:"perPage"`
	Total    int                 `json:"total"`
	Page     int                 `json:"page"`
	Next     *CreatedAtPageQuery `json:"next"`
	Previous *CreatedAtPageQuery `json:"previous"`
}
