package response

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SetMessage(msg string, scss bool) BaseResponse {
	return BaseResponse{
		Message: msg,
		Success: scss,
	}
}
