package utils

type Error struct {
	ErrCode int
	ErrMsg  string
}

func NewError(code int, msg string) *Error {
	return &Error{ErrCode: code, ErrMsg: msg}
}

func NewErrorDefault(msg string) *Error {
	return &Error{ErrCode: 500, ErrMsg: msg}
}

func (err *Error) Error() string {
	return err.ErrMsg
}
func (err *Error) Code() int {
	return err.ErrCode
}
