package apis

import (
	"learn-go/web/core/starter"
	"net/http"
)

type HttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(message string, data interface{}) HttpResponse {
	return HttpResponse{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	}
}
func Fail(code int, message string, data interface{}) HttpResponse {
	// 参数校验异常
	violationError, ok := data.(starter.ConstraintViolationError)
	if ok {
		return HttpResponse{
			Code:    http.StatusBadRequest,
			Message: violationError.Error(),
			Data:    nil,
		}
	}
	// 业务异常
	businessError, ok := data.(starter.BusinessError)
	if ok {
		return HttpResponse{
			Code:    http.StatusBadRequest,
			Message: businessError.Error(),
			Data:    nil,
		}
	}
	return HttpResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
