package response

import (
	"net/http"
)

// Response declare unite adapters response format
type Response struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"msg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// WithData append data into response
func (resp *Response) WithData(data interface{}) *Response {
	return &Response{
		Code:    resp.Code,
		Message: resp.Message,
		Data:    data,
	}
}

var (
	// OK request is success
	OK = &Response{Code: http.StatusOK, Message: "ok"}
)
