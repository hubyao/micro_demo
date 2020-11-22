package errno

import (
	"fmt"
	merror "github.com/micro/go-micro/v2/errors"
	"micro_demo/comm/logging"
)

// Errno ...
type Errno struct {
	Code    int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}

// Err represents an error
type Err struct {
	Code    int
	Message string
	Err     error
}

// New ...
func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Message: errno.Message, Err: err}
}

// Add ...
func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

func (err *Errno) RpcErr() error {
	logging.Logger().Info(merror.New("", err.Message, int32(err.Code)))
	return merror.New("", err.Message, int32(err.Code))
}

func (err *Errno) EqRpcErr(contrast error) bool {
	if contrast.Error() == err.RpcErr().Error() {
		return true
	}
	return false
}

// Addf ...
func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

// Error ...
func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

// DecodeErr ...
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
