package errno

import (
	"fmt"
)

// Errno error info without error
type Errno struct {
	Code    int
	Message string
}

// Err error info include error
type Err struct {
	Code    int
	Message string
	Err     error
}

// New create a new error info
func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Message: errno.Message, Err: err}
}

// Add add a message into the Err's Message
func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

// Addf add a message into the Err's Message by format string and arguments
func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

// IsErrUserNotFound check if the error is a user not found error
func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}

// DecodeErr decode the Err struct
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}

// Error get the error message in Errno
func (err Errno) Error() string {
	return err.Message
}

// Error get the error message in Err
func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}
