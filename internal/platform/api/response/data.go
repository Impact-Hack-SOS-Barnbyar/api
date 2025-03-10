package response

import "net/http"

type DataResponse struct {
	Status int `json:"status"`
	Data   any `json:"data"`
}

func NewDataResponse(data any) *DataResponse {
	return &DataResponse{
		Status: http.StatusOK,
		Data:   data,
	}
}
