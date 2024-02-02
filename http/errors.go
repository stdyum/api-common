package http

import (
	"database/sql"
	"github.com/stdyum/api-common/errors"
	"net/http"
)

var SqlErrorMap = map[error]any{
	sql.ErrTxDone: http.StatusInternalServerError,
	sql.ErrTxDone: http.StatusInternalServerError,
	sql.ErrNoRows: http.StatusNotFound,
}

var ErrorMapper *errors.Mapper

var ErrorMapperBuilder = errors.NewBuilder().
	OnNotFound(http.StatusInternalServerError).
	AppendMap(SqlErrorMap)

func RegisterErrors(errs ...map[error]any) {
	ErrorMapperBuilder.AppendMap(errs...)
	ErrorMapper = ErrorMapperBuilder.Build()
}

func MapError(err error) (int, error) {
	code, err := ErrorMapper.Get(err)
	return code.(int), err
}
