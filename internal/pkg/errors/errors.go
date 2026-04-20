package errors

import "errors"

// 定义错误类型
var (
	ErrNotFound         = errors.New("resource not found")
	ErrInvalidInput     = errors.New("invalid input")
	ErrInternalServer   = errors.New("internal server error")
	ErrDuplicateEntry   = errors.New("duplicate entry")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrForbidden        = errors.New("forbidden")
)

// AppError 应用错误
type AppError struct {
	Err     error
	Message string
	Code    int
}

// Error 实现error接口
func (e *AppError) Error() string {
	return e.Message
}

// NewAppError 创建应用错误
func NewAppError(err error, message string, code int) *AppError {
	return &AppError{
		Err:     err,
		Message: message,
		Code:    code,
	}
}
