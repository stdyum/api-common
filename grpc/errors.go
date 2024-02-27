package grpc

import (
	"database/sql"

	"github.com/stdyum/api-common/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var SqlErrorMap = map[error]any{
	sql.ErrTxDone: codes.Internal,
	sql.ErrTxDone: codes.Internal,
	sql.ErrNoRows: codes.NotFound,
}

var ErrorMapper *errors.Mapper

var ErrorMapperBuilder = errors.NewBuilder().
	OnNotFound(codes.Unknown).
	AppendMap(SqlErrorMap)

func RegisterErrors(errs ...map[error]any) {
	ErrorMapperBuilder.AppendMap(errs...)
	ErrorMapper = ErrorMapperBuilder.Build()
}

func MapError(err error) (codes.Code, error) {
	code, err := ErrorMapper.Get(err)
	return code.(codes.Code), err
}

func ConvertError(err error) error {
	if err == nil {
		return nil
	}

	code, err := MapError(err)
	return status.Error(code, err.Error())
}
