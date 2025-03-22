package utils

import "net/http"

type ResponseError struct {
	Message string                 `json:"message"`
	Status  int                    `json:"message"`
	Error   string                 `json:"error"`
	Data    map[string]interface{} `json:"data"`
}

func BadRequestError(message string, err error, data map[string]interface{}) *ResponseError {
	return &ResponseError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
		Data:    data,
	}
}

func NotFoundRequestError(message string, err error, data map[string]interface{}) *ResponseError {
	return &ResponseError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
		Data:    data,
	}
}

func InternalServerError(message string, err error, data map[string]interface{}) *ResponseError {
	return &ResponseError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server",
		Data:    data,
	}
}
