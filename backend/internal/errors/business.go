package errors

import "net/http"

// BusinessError 是业务异常的统一格式，HTTPStatus 不参与 JSON 序列化。
type BusinessError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HTTPStatus int    `json:"-"`
}

func (e BusinessError) Error() string { return e.Message }

// New 构造一个业务异常。
func New(code, message string, httpStatus int) *BusinessError {
	if httpStatus == 0 {
		httpStatus = http.StatusBadRequest
	}
	return &BusinessError{Code: code, Message: message, HTTPStatus: httpStatus}
}
