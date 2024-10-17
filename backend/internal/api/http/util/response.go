package util

import (
	"github.com/HermanPlay/web-app-backend/internal/api/http/constant"
	"github.com/HermanPlay/web-app-backend/package/domain/schemas"
)

func Null() interface{} {
	return nil
}

func BuildResponse[T any](responseStatus constant.ResponseStatus, data T) schemas.ApiResponse[T] {
	return BuildResponse_(responseStatus.GetResponseStatusCode(), responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data)
}

func BuildResponse_[T any](statuscode int, status, message string, data T) schemas.ApiResponse[T] {
	return schemas.ApiResponse[T]{
		StatusCode:      statuscode,
		ResponseKey:     status,
		ResponseMessage: message,
		Data:            data,
	}
}
