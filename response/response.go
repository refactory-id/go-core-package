package response

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type PaginationResponseData struct {
	CurrentPage    int `json:"current_page"`
	PageCount      int `json:"page_count"`
	PageSize       int `json:"page_size"`
	RowCount       int `json:"total"`
	FirstRowOnPage int `json:"first_row_on_page"`
	LastRowOnPage  int `json:"last_row_on_page"`
}

func SetMessage(msg string, scss bool) BaseResponse {
	return BaseResponse{
		Message: msg,
		Success: scss,
	}
}
