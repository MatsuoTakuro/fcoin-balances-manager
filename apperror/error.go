package apperror

import "errors"

type AppError struct {
	ErrCode     `json:"err_code"`
	ErrMessages []string `json:"err_messages"`
	Err         error    `json:"-"`
}

var _ error = (*AppError)(nil)

func (e *AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// 新規にエラーを作成する場合に使用
func NewAppError(errCode ErrCode, messages ...string) *AppError {
	return &AppError{
		ErrCode:     errCode,
		ErrMessages: messages,
		Err:         errors.New(""),
	}
}
