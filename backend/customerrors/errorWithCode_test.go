package customerrors_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"ygodraft/backend/customerrors"
)

func TestGenericError(t *testing.T) {
	t.Run("test simle error as generic error", func(t *testing.T) {
		// when
		withCodeErr := customerrors.GenericError(assert.AnError)

		// then
		require.Error(t, withCodeErr)
		assert.Contains(t, withCodeErr.Error(), "generic internal server error")
	})
}

func TestWithCode_Error(t *testing.T) {
	t.Run("", func(t *testing.T) {
		// given
		myErr := customerrors.WithCode{
			Params:       []string{"TEST"},
			InternalMsg:  "failed to test %s",
			WrappedError: assert.AnError,
		}

		// when
		errorText := myErr.Error()

		// then
		assert.Equal(t, "failed to test [TEST]: assert.AnError general error for testing", errorText)
	})
}

func TestWithCode_Is(t *testing.T) {
	t.Run("target error is equals with same code", func(t *testing.T) {
		// given
		myErr := customerrors.WithCode{
			Code: "mycode",
		}
		myErr2 := customerrors.WithCode{
			Code: "mycode",
		}

		// when
		isEquals := myErr.Is(myErr2)
		isEquals2 := myErr2.Is(myErr)

		// then
		assert.True(t, isEquals)
		assert.True(t, isEquals2)
	})

	t.Run("target error is not with code error", func(t *testing.T) {
		// given
		myErr := customerrors.WithCode{
			Code: "mycode",
		}

		// when
		isEquals := myErr.Is(assert.AnError)

		// then
		assert.False(t, isEquals)
	})
}

func TestWithCode_Unwrap(t *testing.T) {
	t.Run("test unwrap with wrapped error", func(t *testing.T) {
		// given
		myErr := customerrors.WithCode{
			WrappedError: assert.AnError,
		}

		// when
		err := myErr.Unwrap()

		// then
		assert.Equal(t, assert.AnError, err)
	})
}

func TestWithCode_WithParam(t *testing.T) {
	t.Run("test saving params for the error", func(t *testing.T) {
		// given
		myErr := customerrors.WithCode{InternalMsg: "this is my %s"}

		// when
		myErrWithParam := myErr.WithParam("message")

		// then
		require.Len(t, myErr.Params, 0)
		require.Len(t, myErrWithParam.Params, 1)
		assert.Equal(t, "message", myErrWithParam.Params[0])
	})
}

func TestWithCode_Wrap(t *testing.T) {
	t.Run("test wrap error", func(t *testing.T) {
		// given
		myErr := customerrors.WithCode{}

		// when
		myErrWithWrap := myErr.Wrap(assert.AnError)

		// then
		assert.Nil(t, myErr.WrappedError)
		assert.Equal(t, assert.AnError, myErrWithWrap.WrappedError)
	})
}
