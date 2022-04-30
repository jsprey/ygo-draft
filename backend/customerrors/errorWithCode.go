package customerrors

import (
	"fmt"
)

var errGenericInternal = WithCode{Code: CodeGenericInternalError, InternalMsg: "generic internal server error"}

// WithCode is an error which contains an error code
type WithCode struct {
	Code         string
	Params       []string
	InternalMsg  string
	WrappedError error
}

// Is checks whether the target error is the same by comparing the error code
func (w WithCode) Is(target error) bool {
	t, ok := target.(WithCode)
	if !ok {
		return false
	}
	return w.Code == t.Code
}

// Error builds a string from the error
func (w WithCode) Error() string {
	internalMessage := w.InternalMsg
	if w.Params != nil {
		internalMessage = fmt.Sprintf(w.InternalMsg, w.Params)
	}

	if w.WrappedError != nil {
		return fmt.Sprintf("%s: %s", internalMessage, w.WrappedError)
	}
	return internalMessage
}

// Unwrap is used to make WithCode work with errors.Is and errors.As.
func (w WithCode) Unwrap() error {
	return w.WrappedError
}

// Wrap is used to wrap an error with a WithCode err
func (w WithCode) Wrap(err error) WithCode {
	w.WrappedError = err
	return w
}

// WithParam is used to return the error with params attached
func (w WithCode) WithParam(params ...string) WithCode {
	newError := w
	newError.Params = params
	return newError
}

// GenericError is used to generate an error with a user friendly message for technical problems
func GenericError(err error) WithCode {
	genericError := errGenericInternal
	genericError.WrappedError = err
	return genericError
}
