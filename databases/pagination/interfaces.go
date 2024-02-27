package pagination

type Query interface {
	Field() string
	Order() string
	Cursor() any
	PerPage() int
	Page() int
}
