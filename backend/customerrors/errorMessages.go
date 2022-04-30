package customerrors

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// general errors
const (
	CodeGenericInternalError = "CODE_GENERIC_INTERNAL_ERROR"
	CodeCardDoesNotExist     = "CODE_CARD_DOES_NOT_EXIST"
)

var germanMessages = map[string]string{
	// general errors
	CodeGenericInternalError: "Ein interner Serverfehler ist aufgetreten.",
	CodeCardDoesNotExist:     "Die Karte [%s] existiert nicht.",
}

type errorResponse struct {
	Message string `json:"message"`
}

// WriteErrorWithMessage builds a message from the given errors and writes the error in the gin context
func WriteErrorWithMessage(c *gin.Context, err WithCode, embeddedError error, statusCode int) {
	err.WrappedError = embeddedError
	_ = c.Error(err)

	errorMessage := BuildErrorMessage(err)
	res := errorResponse{Message: errorMessage}

	c.JSON(statusCode, res)
}

func BuildErrorMessage(err WithCode) string {
	messages := []string{}
	var inspectedError error = err
	// Checks whether the inspectedError contains an errorCode and appends its message. Then unwraps the inspectedError and
	// repeats the process until there is no more errorCode.
	for {
		var errorWithCode WithCode
		if errors.As(inspectedError, &errorWithCode) {
			params := make([]interface{}, len(errorWithCode.Params))
			for i, v := range errorWithCode.Params {
				params[i] = v
			}
			messages = append(messages, fmt.Sprintf(germanMessages[errorWithCode.Code], params...))
			inspectedError = errorWithCode.Unwrap()
		} else {
			break
		}
	}

	return strings.Join(messages, " ")
}
