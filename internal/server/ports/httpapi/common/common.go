package common

const (
	HeaderAuthorization = "Authorization"
	AuthTypeTelegram    = "tma"
	AuthTypeDebug       = "dbg"
)

// Status defines model for Status.
type Status struct {
	// Code Error code
	Code int `json:"code"`

	// Message Error message
	Message string `json:"message"`
}

func NewStatusResponse(code int, message string) Status {
	return Status{
		Code:    code,
		Message: message,
	}
}
