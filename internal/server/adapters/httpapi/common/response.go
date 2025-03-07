package common

type Status struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type StatusResponse struct {
	Status Status `json:"status"`
}

func NewStatusResponse(code int64, message string) StatusResponse {
	return StatusResponse{
		Status: Status{
			Code:    code,
			Message: message,
		},
	}
}
