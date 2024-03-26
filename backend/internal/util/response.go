package util

import "github.com/gin-gonic/gin"

const (
	MessageError   = 0
	MessageSuccess = 1
)

func MakeMessage(messType int, message string, data interface{}) gin.H {
	response := gin.H{}

	switch messType {
	case MessageError:
		response["status"] = "error"
	case MessageSuccess:
		response["status"] = "success"
	default:
		response["status"] = "success"
	}

	if message != "" {
		response["message"] = message
	}

	if data != nil {
		response["data"] = data
	}

	return response
}
