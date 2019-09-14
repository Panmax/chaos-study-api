package common

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Pagination struct {
	Total   uint32      `json:"total"`
	Results interface{} `json:"results"`
}

func NewSuccessResponse(data interface{}) Response {
	response := Response{Code: 0, Message: "success", Data: data}
	return response
}
