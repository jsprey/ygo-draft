package customerrors_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"ygodraft/backend/customerrors"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

var (
	ErrorInternal1 = customerrors.WithCode{
		Code:        "CODE_TEST_INTERNAL_1",
		InternalMsg: "internal 1",
	}
	ErrorInternal2 = customerrors.WithCode{
		Code:        "CODE_TEST_INTERNAL_2",
		InternalMsg: "internal 2",
	}
	ErrorInternal3 = customerrors.WithCode{
		Code:        "CODE_TEST_INTERNAL_3",
		InternalMsg: "internal 3",
	}
	ErrorInternal4 = customerrors.WithCode{
		Code:        "CODE_TEST_INTERNAL_4",
		InternalMsg: "internal 4",
	}
	ErrorOK = customerrors.WithCode{
		Code:        "CODE_OK",
		InternalMsg: "ok error",
	}
)

var statusCodeMap = map[string]int{
	ErrorInternal1.Code: http.StatusBadRequest,
	ErrorInternal2.Code: http.StatusUnauthorized,
	ErrorInternal3.Code: http.StatusNotFound,
	ErrorOK.Code:        http.StatusOK,
}

func TestHandleError(t *testing.T) {
	tests := []struct {
		name               string
		internalError      error
		expectedStatusCode int
		expectError        bool
	}{
		{name: "internal 1 => StatusBadRequest", internalError: ErrorInternal1, expectedStatusCode: http.StatusBadRequest, expectError: true},
		{name: "internal 2 => StatusUnauthorized", internalError: ErrorInternal2, expectedStatusCode: http.StatusUnauthorized, expectError: true},
		{name: "internal 3 => StatusNotFound", internalError: ErrorInternal3, expectedStatusCode: http.StatusNotFound, expectError: true},
		{name: "internal 4 => StatusInternalServerError", internalError: ErrorInternal4, expectedStatusCode: http.StatusInternalServerError, expectError: true},
		{name: "nil error does return write an error into the context and returns false", internalError: nil, expectError: false},
		{name: "map containing a 200 does return write an error into the context and returns false", internalError: ErrorOK, expectError: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := gin.CreateTestContext(httptest.NewRecorder())

			contextError := customerrors.WithCode{
				Code:        "test",
				InternalMsg: "test",
			}
			customerrors.HandleErrorAndShouldAbort(c, statusCodeMap, contextError, tt.internalError)

			if tt.expectError {
				assert.Equal(t, tt.expectedStatusCode, c.Writer.Status())
			}
		})
	}
}
