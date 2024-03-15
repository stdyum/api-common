package errors

type CodeError struct {
	Err  error
	Code any
}
