package response

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SetMessage(message string, success bool) BaseResponse {
	return BaseResponse{
		Message: message,
		Success: success,
	}
}
