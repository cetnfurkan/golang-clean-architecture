package errors

import "github.com/cetnfurkan/core/errors"

type ErrorCode struct {
	Code    errors.HttpErrorCode
	Message string
}

var (
	DATABASE_ERROR     = ErrorCode{Code: "1", Message: "db_error"}
	GENERIC_ERROR      = ErrorCode{Code: "2", Message: "generic_error"}
	GRPC_ERROR         = ErrorCode{Code: "3", Message: "grpc_error"}
	INVALID_PARAMETERS = ErrorCode{Code: "4", Message: "invalid_parameters"}
)
