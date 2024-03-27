package util

const (
	MessageError   = 0
	MessageSuccess = 1
)

type MessageResponse map[string]any

func MakeMessage(messType int, message string, data interface{}) MessageResponse {
	response := MessageResponse{}

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
