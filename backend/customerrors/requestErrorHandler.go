package customerrors

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleErrorAndShouldAbort determines the correct status code of an error regarding the current context and returns
// true when the requestHandler should abort the handling of the request.
func HandleErrorAndShouldAbort(c *gin.Context, statusCodeMap map[string]int, contextError WithCode, internalError error) bool {
	if internalError == nil {
		return false
	}

	var internalErrorWithCode WithCode
	if errors.As(internalError, &internalErrorWithCode) {
		code, ok := statusCodeMap[internalErrorWithCode.Code]
		if ok && code == http.StatusOK {
			return false
		} else if ok {
			WriteErrorWithMessage(c, contextError, internalError, code)
			return true
		} else {
			WriteErrorWithMessage(c, contextError, internalError, http.StatusInternalServerError)
			return true
		}
	} else {
		WriteErrorWithMessage(c, contextError, GenericError(internalError), http.StatusInternalServerError)
		return true
	}
}
