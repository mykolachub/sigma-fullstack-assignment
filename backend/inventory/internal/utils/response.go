package utils

type Response struct {
	// Code    int                    `json:"code"`
	Message string                 `json:"message,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Status  string                 `json:"status"`
}

func NewResponse(message string /* code int */) *Response {
	return &Response{
		Status:  "success",
		Message: message,
		// Code:    code,
		Data: map[string]interface{}{},
	}
}

func NewErrResponse(message string /* code int */) Response {
	return Response{
		Status:  "error",
		Message: message,
		// Code:    code,
	}
}

func (r *Response) AddKey(key string, value interface{}) *Response {
	r.Data[key] = value
	return r
}
