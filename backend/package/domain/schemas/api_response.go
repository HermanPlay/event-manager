package schemas

type ApiResponse[T any] struct {
	StatusCode      int    `json:"status_code"`
	ResponseKey     string `json:"response_key"`
	ResponseMessage string `json:"response_message"`
	Data            T      `json:"data"`
}
