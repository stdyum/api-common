package http

import (
	"database/sql"
	"errors"
	"net/http"

	errorsMapper "github.com/stdyum/api-common/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var SqlErrorMap = map[error]any{
	sql.ErrTxDone: http.StatusInternalServerError,
	sql.ErrTxDone: http.StatusInternalServerError,
	sql.ErrNoRows: http.StatusNotFound,
}

var CustomErrorMap = map[error]any{
	errorsMapper.ErrValidation:   http.StatusUnprocessableEntity,
	errorsMapper.ErrUnauthorized: http.StatusUnauthorized,
}

var ErrorMapper *errorsMapper.Mapper

var ErrorMapperBuilder = errorsMapper.NewBuilder().
	OnNotFound(http.StatusInternalServerError).
	AppendMap(CustomErrorMap).
	AppendMap(SqlErrorMap)

func RegisterErrors(errs ...map[error]any) {
	ErrorMapperBuilder.AppendMap(errs...)
	ErrorMapper = ErrorMapperBuilder.Build()
}

func MapError(err error) (int, error) {
	st, ok := status.FromError(errors.Unwrap(err))
	if ok {
		return MapGrpcErrors(st), errors.New(st.Message())
	}

	code, err := ErrorMapper.Get(err)
	return code.(int), err
}

func MapGrpcErrors(s *status.Status) int {
	switch s.Code() {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		return http.StatusPreconditionFailed
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
