package customerrors_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"ygodraft/backend/customerrors"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

var simpleErrorWithCode = customerrors.WithCode{
	Code:         customerrors.CodeGenericInternalError,
	Params:       nil,
	InternalMsg:  "test error",
	WrappedError: nil,
}

var embeddingError = customerrors.WithCode{
	Code:         customerrors.CodeCardDoesNotExist,
	Params:       []string{"115"},
	InternalMsg:  "embedding error",
	WrappedError: nil,
}

var errorWithParams = customerrors.WithCode{
	Code:         customerrors.CodeCardDoesNotExist,
	Params:       []string{"115"},
	InternalMsg:  "test error '%s'",
	WrappedError: nil,
}

var internalWithParams = customerrors.WithCode{
	Code:         customerrors.CodeCardDoesNotExist,
	Params:       []string{"115"},
	InternalMsg:  "internal error with params '%s'",
	WrappedError: nil,
}

func TestWriteErrorWithMessage_SimpleErrorWithCode(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	customerrors.WriteErrorWithMessage(c, simpleErrorWithCode, nil, http.StatusBadRequest)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	body, _ := io.ReadAll(w.Body)
	assert.Equal(t, "{\"message\":\"Ein interner Serverfehler ist aufgetreten.\"}", string(body))
	assert.Equal(t, "Error #01: test error\n", c.Errors.String())
}

func TestWriteErrorWithMessage_WithParams(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	customerrors.WriteErrorWithMessage(c, errorWithParams, nil, http.StatusNotFound)

	assert.Equal(t, http.StatusNotFound, w.Code)
	body, _ := io.ReadAll(w.Body)
	assert.Equal(t, "{\"message\":\"Die Karte [115] existiert nicht.\"}", string(body))
}

func TestWriteErrorWithMessage_MultiError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	internalError := fmt.Errorf("internal error: %w", simpleErrorWithCode)

	customerrors.WriteErrorWithMessage(c, embeddingError, internalError, http.StatusBadRequest)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	body, _ := io.ReadAll(w.Body)
	assert.Equal(t, "{\"message\":\"Die Karte [115] existiert nicht. Ein interner Serverfehler ist aufgetreten.\"}", string(body))
}

func TestWriteErrorWithMessage_MultiWithParams(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	customerrors.WriteErrorWithMessage(c, errorWithParams, internalWithParams, http.StatusNotFound)

	assert.Equal(t, http.StatusNotFound, w.Code)
	body, _ := io.ReadAll(w.Body)
	assert.Equal(t, "{\"message\":\"Die Karte [115] existiert nicht. Die Karte [115] existiert nicht.\"}", string(body))
}
