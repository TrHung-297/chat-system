package response

import "github.com/TrHung-297/chat-v2/herror"

type Response struct {
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
}

// ErrorResponse defines an herror response object
type ErrorResponse struct {
	ErrorCode uint32 `json:"ErrorCode"`
	Message   string `json:"Message"`
	Exception string `json:"Exception"`
}

func NewGerrorResponse(gerr *herror.Error, userId string) (string, ErrorResponse) {
	msg := herror.T(gerr.Code)
	error := ErrorResponse{
		ErrorCode: gerr.Code,
		Message:   "[BACKEND]" + gerr.Error.Error(),
		Exception: gerr.Line + ":[" + userId + "]",
	}
	return msg, error
}

// NewErrorResponse creates a new herror response object
func NewErrorResponse(errorCode uint32, message string, exception string) (string, ErrorResponse) {
	msg := herror.T(errorCode)
	error := ErrorResponse{
		ErrorCode: errorCode,
		Message:   "[BACKEND]" + message,
		Exception: exception,
	}
	return msg, error
}