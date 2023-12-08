package error

type ErrCode int

type Error struct {
	msg  string
	code ErrCode
}

// Error returns error message
func (e *Error) Error() string {
	return e.msg
}

// Code returns code of error
func (e *Error) Code() ErrCode {
	return e.code
}

func NewError(msg error, code ErrCode) *Error {
	return &Error{msg.Error(), code}
}
