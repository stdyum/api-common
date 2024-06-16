package pagination

type Query interface {
	Field() string
	SetField(field string)

	Search() string
	Order() string
	Cursor() any
	PerPage() int
	Page() int
}
