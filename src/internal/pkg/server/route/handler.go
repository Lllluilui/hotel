package route

import (
	"net/http"
)

// HandlerResult 处理结果
type HandlerResult struct {
	// Status 状态码
	Status int `json:"status"`

	// Message 提示信息
	Message string `json:"message"`

	// Data 数据
	Data any `json:"data"`
}

// NewHandlerSuccess
//
//	@desc: 处理正确结果
//	@param data any
//	@return *HandlerResult
func NewHandlerSuccess(data any) *HandlerResult {
	return NewHandlerResult(http.StatusOK, "", data)
}

// NewHandlerError
//
//	@desc: 处理错误
//	@param message string
//	@param data any
//	@param statuses ...int
//	@return *HandlerResult
func NewHandlerError(message string, data any, statuses ...int) *HandlerResult {
	// 默认是内部错误
	status := http.StatusInternalServerError

	if len(statuses) > 0 {
		status = statuses[0]
	}

	return NewHandlerResult(status, message, data)
}

// NewHandlerResult
//
//	@desc: 处理结果
//	@param status int
//	@param message string
//	@param data any
//	@return *HandlerResult
func NewHandlerResult(status int, message string, data any) *HandlerResult {

	return &HandlerResult{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
